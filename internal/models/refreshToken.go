package models

import "time"

type RefreshToken struct {
	ID         int        `gorm:"primarykey" json:"id"`
	Token      string     `gorm:"size:255;not null" json:"token"`
	UserID     uint       `gorm:"not null" json:"user_id"`
	ExpiresAt  time.Time  `gorm:"not null" json:"expires_at"`
	Revoked_at *time.Time `gorm:"default:null" json:"revoked_at"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	User User `gorm:"foreignKey:UserID" json:"-"`
}

func (rt *RefreshToken) TableName() string {
	return "refesh_token"
}
