package handlers

import (
	"net/http"
	"strconv"

	"github.com/cbitbaly/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetMyInfoHandler(c *gin.Context) {
	userId, err := c.Get("userID")

	if !err {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  "Vous n'êtes pas connecté",
			"status": http.StatusUnauthorized,
		})
		return
	}

	user, erreur := h.service.GetUserWithPosts(userId.(uint))

	if erreur != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  erreur.Error(),
			"status": http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"status": http.StatusOK,
	})

}

func (h *UserHandler) GetAllUserHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	users, total, err := h.service.GetAllUsers(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users":  users,
		"total":  total,
		"status": http.StatusOK,
	})
}

func (h *UserHandler) GetUserByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	user, erreur := h.service.GetUserByID(uint(id))

	if erreur != nil {

		if erreur.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error":  erreur.Error(),
				"status": http.StatusNotFound,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  erreur.Error(),
			"status": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"status": http.StatusOK,
	})
}

func (h *UserHandler) UpdateUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	var data map[string]any

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	erreur := h.service.UpdateUser(uint(id), data)

	if erreur != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   erreur.Error(),
			"message": "impossible de mettre à jour l'utilisateur",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "utilisateur mis à jour avec succès",
		"status":  http.StatusOK,
	})
}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	erreur := h.service.DeleteUser(uint(id))

	if erreur != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   erreur.Error(),
			"message": "impossible de supprimer l'utilisateur",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "utilisateur supprimé avec succès",
		"status":  http.StatusOK,
	})
}
