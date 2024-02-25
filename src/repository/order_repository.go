package repository

import (
	"reku-technical-test/src/entity"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAll() ([]entity.Order, error)
	FindByID(id int) (entity.Order, error)
	Save(order entity.Order) (entity.Order, error)
	UpdateStatus(id int) error
}

type orderRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewOrder(db *gorm.DB, redis *redis.Client) *orderRepository {
	return &orderRepository{db, redis}
}

func (r *orderRepository) FindAll() ([]entity.Order, error) {
	var order []entity.Order

	err := r.db.
		Find(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *orderRepository) FindByID(id int) (entity.Order, error) {
	var order entity.Order

	err := r.db.
		Where("id = ?", id).
		Find(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *orderRepository) UpdateStatus(id int) error {
	var order entity.Order

	err := r.db.Where("id = ?", id).
		First(&order).Error
	if err != nil {
		return err
	}

	order.Status = "ready"
	r.db.Save(&order)

	return nil
}

func (r *orderRepository) Save(order entity.Order) (entity.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}
