package repositories

import (
	"github.com/cbitbaly/internal/models"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r *RefreshTokenRepository) Create(token *models.RefreshToken) error {
	return r.db.Create(&token).Error
}

func (r *RefreshTokenRepository) FindByToken(token string) (*models.RefreshToken, error) {
	var RefreshToken models.RefreshToken

	err := r.db.Preload("User").Where("token = ?", token).First(&RefreshToken).Error

	if err != nil {
		return nil, err
	}
	return &RefreshToken, nil
}

func (r *RefreshTokenRepository) Update(id uint, data map[string]any) error {
	return r.db.Model(&models.RefreshToken{}).Updates(data).Error
}

func (r *RefreshTokenRepository) Delete(id uint) error {
	return r.db.Delete(&models.RefreshToken{}, id).Error
}
