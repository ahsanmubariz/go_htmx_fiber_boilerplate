package database

import (
	"log"

	"github.com/ahsanmubariz/go_htmx_fiber_boilerplate/internal/config"
	"github.com/ahsanmubariz/go_htmx_fiber_boilerplate/internal/modules/users"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Database connection established.")
	return db, nil
}

func RunMigrations(db *gorm.DB) error {
	log.Println("Running database migrations...")
	err := db.AutoMigrate(&users.User{})
	if err != nil {
		return err
	}
	log.Println("Database migrations completed.")
	return nil
}
