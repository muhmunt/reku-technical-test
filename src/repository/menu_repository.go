package repository

import (
	"reku-technical-test/src/entity"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type MenuRepository interface {
	FindAll() ([]entity.Menu, error)
	FindByID(id int) (entity.Menu, error)
}

type menuRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewMenu(db *gorm.DB, redis *redis.Client) *menuRepository {
	return &menuRepository{db, redis}
}

func (r *menuRepository) FindAll() ([]entity.Menu, error) {
	var menu []entity.Menu

	err := r.db.
		Find(&menu).Error

	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (r *menuRepository) FindByID(id int) (entity.Menu, error) {
	var menu entity.Menu

	err := r.db.
		Where("id = ?", id).
		First(&menu).Error

	if err != nil {
		return menu, err
	}

	return menu, nil
}
