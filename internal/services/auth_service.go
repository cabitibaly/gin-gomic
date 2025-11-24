package services

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/cbitbaly/internal/models"
	"github.com/cbitbaly/internal/repositories"
	"github.com/cbitbaly/pkg/utils"
)

type AuthService struct {
	jwtRepo  *repositories.JWTRepository
	userRepo *repositories.UserRepository
}

func NewAuthService(jwtRepo *repositories.JWTRepository, userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		jwtRepo:  jwtRepo,
		userRepo: userRepo,
	}
}

func (s *AuthService) RegisterUser(nom, prenom, email, password string) (*models.User, error) {
	if nom == "" || email == "" || password == "" {
		return nil, fmt.Errorf("nom, email et mot de passe sont obligatoires")
	}

	userExist, _ := s.userRepo.FindByEmail(email)

	if userExist != nil {
		return nil, fmt.Errorf("email déjà utilisé")
	}

	hashedPassword, err := utils.HashPassword(password)

	if err != nil {
		return nil, errors.New("une erreur est survenue lors du hashage du mot de passe")
	}

	user := &models.User{
		Nom:      nom,
		Prenom:   prenom,
		Email:    email,
		Password: hashedPassword,
	}

	erreur := s.userRepo.Create(user)

	if erreur != nil {
		return nil, erreur
	}

	return user, nil
}

func (s *AuthService) LoginUser(email, password string) (*models.User, string, error) {
	if email == "" || password == "" {
		return nil, "", fmt.Errorf("email et mot de passe sont obligatoires")
	}

	userExist, _ := s.userRepo.FindByEmail(email)

	if userExist == nil {
		return nil, "", fmt.Errorf("email inconnu")
	}

	if !utils.ComparePassword(password, userExist.Password) {
		return nil, "", fmt.Errorf("email ou mot de passe incorrect")
	}

	tokenString, err := utils.GenerateToken(uint(userExist.ID), userExist.Email)

	if err != nil {
		log.Println("Erreur lors de la génération du token :", err)
		return nil, "", errors.New("erreur lors de la génération du token")
	}

	token := &models.Jwt{
		Token:     tokenString,
		UserID:    uint(userExist.ID),
		ExpiresAt: time.Now().Add(time.Hour * 72),
	}

	err = s.jwtRepo.Create(token)

	if err != nil {
		return nil, "", err
	}

	return userExist, tokenString, nil
}
