package handlers

import (
	"net/http"

	"github.com/cbitbaly/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) RegisterUserHandler(c *gin.Context) {
	var data struct {
		Nom      string `json:"nom"`
		Prenom   string `json:"prenom"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err,
			"status": http.StatusBadRequest,
		})
		return
	}

	user, err := h.service.RegisterUser(data.Nom, data.Prenom, data.Email, data.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":   user,
		"status": http.StatusCreated,
	})
}

func (h *AuthHandler) LoginUserHandler(c *gin.Context) {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	user, token, err := h.service.LoginUser(data.Email, data.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  err.Error(),
			"status": http.StatusUnauthorized,
		})
		return
	}

	c.SetCookie(
		"jwt",
		token,
		3600*24*3,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"status": http.StatusOK,
	})
}
