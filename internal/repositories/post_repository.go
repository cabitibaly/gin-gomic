package repositories

import (
	"fmt"

	"github.com/cbitbaly/internal/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *models.Post) error {
	return r.db.Create(&post).Error
}

func (r *PostRepository) FindById(id uint) (*models.Post, error) {
	var post models.Post

	err := r.db.First(&post, id).Error

	if err != nil {
		return nil, fmt.Errorf("post non trouvé")
	}

	return &post, nil
}

func (r *PostRepository) FindAll(page, pageSize int) ([]models.Post, int64, error) {
	var posts []models.Post
	var total int64

	r.db.Model(&models.Post{}).Count(&total)

	offset := (page - 1) * pageSize
	err := r.db.Offset(offset).Limit(pageSize).Find(&posts).Error

	return posts, total, err
}

func (r *PostRepository) FindByUserId(userID uint) ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Where("user_id = ?", userID).Find(&posts).Error

	if err != nil {
		return nil, fmt.Errorf("aucun post trouvé")
	}

	return posts, nil
}

func (r *PostRepository) Update(id uint, data map[string]any) error {
	return r.db.Model(&models.Post{}).Where("id = ?", id).Updates(data).Error
}

func (r *PostRepository) Delete(id uint) error {
	return r.db.Delete(&models.Post{}, id).Error
}
