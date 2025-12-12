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

	user, refreshToken, accessToken, err := h.service.LoginUser(data.Email, data.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":          user,
		"refresh_token": refreshToken,
		"access_token":  accessToken,
		"status":        http.StatusOK,
	})
}

func (h *AuthHandler) RefreshTokenHandler(c *gin.Context) {
	var data struct {
		Token string `json:"token"`
	}

	if err := c.ShouldBindJSON((&data)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	refreshToken, accessToken, err := h.service.NouveauRefreshToken(data.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"refresh_token": refreshToken,
		"access_token":  accessToken,
		"status":        http.StatusOK,
	})

}
