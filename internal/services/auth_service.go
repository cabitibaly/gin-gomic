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
	refreshTokenRepo *repositories.RefreshTokenRepository
	userRepo         *repositories.UserRepository
}

func NewAuthService(refreshTokenRepo *repositories.RefreshTokenRepository, userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		refreshTokenRepo: refreshTokenRepo,
		userRepo:         userRepo,
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

func (s *AuthService) LoginUser(email, password string) (*models.User, string, string, error) {
	if email == "" || password == "" {
		return nil, "", "", fmt.Errorf("email et mot de passe sont obligatoires")
	}

	userExist, _ := s.userRepo.FindByEmail(email)

	if userExist == nil {
		return nil, "", "", fmt.Errorf("email inconnu")
	}

	if !utils.ComparePassword(password, userExist.Password) {
		return nil, "", "", fmt.Errorf("email ou mot de passe incorrect")
	}

	accessToken, errAT := utils.GenerateAccessToken(uint(userExist.ID), userExist.Email)
	if errAT != nil {
		log.Println("Une erreur est survenue lors de la génération du token d'accès :", errAT)
		return nil, "", "", fmt.Errorf("une erreur est survenue lors de la génération du token d'accès")
	}

	refreshToken, expireTime, errRT := utils.GenerateRefreshToken(uint(userExist.ID), userExist.Email)
	if errRT != nil {
		log.Println("Une erreur est survenue lors de la génération du refresh token :", errRT)
		return nil, "", "", fmt.Errorf("une erreur est survenue lors de la génération du refresh token")
	}

	token := &models.RefreshToken{
		Token:     refreshToken,
		UserID:    uint(userExist.ID),
		ExpiresAt: *expireTime,
	}

	err := s.refreshTokenRepo.Create(token)

	if err != nil {
		return nil, "", "", err
	}

	return userExist, refreshToken, accessToken, nil
}

func (s *AuthService) NouveauRefreshToken(refreshToken string) (string, string, error) {
	ancienRT, err := s.refreshTokenRepo.FindByToken(refreshToken)

	if err != nil {
		return "", "", err
	}

	location, _ := time.LoadLocation("Africa/Ouagadougou")
	maintenat := time.Now().In(location)
	if maintenat.After(ancienRT.ExpiresAt) || ancienRT.Revoked_at != nil {
		return "", "", fmt.Errorf("le refresh token a expiré ou a été révoqué")
	}

	accessToken, errAT := utils.GenerateAccessToken(ancienRT.UserID, ancienRT.User.Email)
	if errAT != nil {
		log.Println("Une erreur est survenue lors de la génération du token d'accès :", errAT)
		return "", "", fmt.Errorf("une erreur est survenue lors de la génération du token d'accès")
	}

	newRefreshToken, expireTime, errRT := utils.GenerateRefreshToken(ancienRT.UserID, ancienRT.User.Email)
	if errRT != nil {
		log.Println("Une erreur est survenue lors de la génération du refresh token :", errRT)
		return "", "", fmt.Errorf("une erreur est survenue lors de la génération du refresh token")
	}

	s.refreshTokenRepo.Delete(uint(ancienRT.ID))

	token := &models.RefreshToken{
		Token:     newRefreshToken,
		ExpiresAt: *expireTime,
		UserID:    ancienRT.UserID,
	}

	err = s.refreshTokenRepo.Create(token)

	if err != nil {
		return "", "", err
	}

	return newRefreshToken, accessToken, nil
}
