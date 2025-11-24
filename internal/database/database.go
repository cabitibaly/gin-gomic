package database

import (
	"fmt"
	"log"
	"time"

	"github.com/cbitbaly/config"
	"github.com/cbitbaly/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("impossible de se connecter à la base de donnée")
	}

	sqlDB, erreur := DB.DB()

	if erreur != nil {
		return erreur
	}

	sqlDB.SetConnMaxIdleTime(time.Hour)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)

	log.Println("Connexion à la base de données réussie....✅")

	return nil

}

func GetDB() *gorm.DB {
	return DB
}

func Migration() error {

	err := DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Jwt{},
	)

	if err != nil {
		return fmt.Errorf("une erreur est survenue lors de la migration %s", err)
	}

	log.Println("Migrations exécutées avec succès....✅")

	return nil
}
