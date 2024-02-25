package repository

import (
	"reku-technical-test/src/entity"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ChefRepository interface {
	FindAvailable() (entity.Chef, error)
	Save(chef entity.Chef) (entity.Chef, error)
	FindByID(id int) (entity.Chef, error)
	UpdateStatus(id, status int) error
}

type chefRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewChef(db *gorm.DB, redis *redis.Client) *chefRepository {
	return &chefRepository{db, redis}
}

func (r *chefRepository) Save(chef entity.Chef) (entity.Chef, error) {
	err := r.db.Create(&chef).Error
	if err != nil {
		return chef, err
	}

	return chef, nil
}

func (r *chefRepository) FindAvailable() (entity.Chef, error) {
	var chef entity.Chef

	err := r.db.
		Where("busy = ?", false).
		First(&chef).Error

	if err != nil {
		return chef, err
	}

	return chef, nil
}

func (r *chefRepository) FindByID(id int) (entity.Chef, error) {
	var chef entity.Chef

	err := r.db.
		Where("id = ?", id).
		First(&chef).Error

	if err != nil {
		return chef, err
	}

	return chef, nil
}

func (r *chefRepository) UpdateStatus(id, status int) error {
	var chef entity.Chef

	err := r.db.Where("id = ?", id).
		First(&chef).Error
	if err != nil {
		return err
	}

	chef.Busy = status
	r.db.Save(&chef)

	return nil
}
