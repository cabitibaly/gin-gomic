package models

import "time"

type User struct {
	ID           int            `gorm:"primarykey" json:"id"`
	Nom          string         `gorm:"size:20;not null" json:"nom"`
	Prenom       string         `gorm:"size:20;" json:"prenom"`
	Email        string         `gorm:"size:100;not null;uniqueIndex" json:"email"`
	Password     string         `gorm:"size:100;not null" json:"-"`
	CreateAt     time.Time      `gorm:"autoCreateTime" json:"create_at"`
	UpdateAt     time.Time      `gorm:"autoUpdateTime" json:"update_at"`
	Posts        []Post         `gorm:"foreignKey:UserID" json:"posts,omitempty"`
	RefreshToken []RefreshToken `gorm:"foreignKey:UserID" json:"refresh_token,omitempty"`
}

func (u *User) TableName() string {
	return "user"
}
