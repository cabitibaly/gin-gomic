package services

import (
	"fmt"

	"github.com/cbitbaly/internal/models"
	"github.com/cbitbaly/internal/repositories"
)

type PostService struct {
	repo *repositories.PostRepository
}

func NewPostService(repo *repositories.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(userID uint, title, content string) error {
	if title == "" || content == "" {
		return fmt.Errorf("le titre et le contenu ne peuvent pas être vides")
	}

	post := models.Post{
		UserID:  userID,
		Title:   title,
		Content: content,
	}

	return s.repo.Create(&post)
}

func (s *PostService) GetPostById(postID uint) (*models.Post, error) {
	return s.repo.FindById(postID)
}

func (s *PostService) GetAllPosts(page, pageSize int) ([]models.Post, int64, error) {
	return s.repo.FindAll(page, pageSize)
}

func (s *PostService) GetAllPostsByUserID(userID uint) ([]models.Post, error) {
	return s.repo.FindByUserId(userID)
}

func (s *PostService) UpdatePost(postID uint, data map[string]any) error {
	return s.repo.Update(postID, data)
}

func (s *PostService) DeletePost(postID uint) error {
	return s.repo.Delete(postID)
}
