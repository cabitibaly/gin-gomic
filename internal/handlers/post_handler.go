package handlers

import (
	"net/http"
	"strconv"

	"github.com/cbitbaly/internal/services"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service *services.PostService
}

func NewPostHandler(service *services.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) CreatePostHandler(c *gin.Context) {
	userID, err := c.Get("userID")

	if !err {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  "Vous n'êtes pas connecté",
			"status": http.StatusUnauthorized,
		})
		return
	}

	var data struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	erreur := h.service.CreatePost(userID.(uint), data.Title, data.Content)

	if erreur != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   erreur.Error(),
			"message": "impossible de créer le post",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "post créé avec succès",
		"status":  http.StatusOK,
	})
}

func (h *PostHandler) GetPostByIdHandler(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	post, erreur := h.service.GetPostById(uint(postID))

	if erreur != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   erreur.Error(),
			"message": "impossible de récupérer le post",
			"status":  http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post":   post,
		"status": http.StatusOK,
	})
}

func (h *PostHandler) GetAllPostsHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	posts, total, err := h.service.GetAllPosts(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err,
			"status": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts":  posts,
		"total":  total,
		"status": http.StatusOK,
	})
}

func (h *PostHandler) GetAllPostsByUserIDHandler(c *gin.Context) {
	userId, err := c.Get("userID")

	if !err {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  "Vous n'êtes pas connecté",
			"status": http.StatusUnauthorized,
		})
		return
	}

	posts, erreur := h.service.GetAllPostsByUserID(userId.(uint))

	if erreur != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   erreur.Error(),
			"message": "impossible de récupérer les posts",
			"status":  http.StatusNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts":  posts,
		"status": http.StatusOK,
	})
}

func (h *PostHandler) UpdatePostHandler(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))

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

	if data["title"] == "" || data["content"] == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "le titre et le contenu ne peuvent pas être vides",
			"status": http.StatusBadRequest,
		})
		return
	}

	_, errExist := h.service.GetPostById(uint(postID))

	if errExist != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  errExist.Error(),
			"status": http.StatusNotFound,
		})
		return
	}

	erreur := h.service.UpdatePost(uint(postID), data)

	if erreur != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   erreur.Error(),
			"message": "impossible de mettre à jour le post",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "post mis à jour avec succès",
		"status":  http.StatusOK,
	})
}

func (h *PostHandler) DeletePostHandler(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	_, errExist := h.service.GetPostById(uint(postID))

	if errExist != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  errExist.Error(),
			"status": http.StatusNotFound,
		})
		return
	}

	erreur := h.service.DeletePost(uint(postID))

	if erreur != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   erreur.Error(),
			"message": "impossible de supprimer le post",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "post supprimé avec succès",
		"status":  http.StatusOK,
	})
}
