package repository

import (
	"log"
	"reku-technical-test/src/entity"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ShortenerRepository interface {
	Save(shortener entity.Shortener) (entity.Shortener, error)
	FindByID(url string) (entity.Shortener, error)
	RedisSave(c *gin.Context, shortener entity.Shortener) (entity.Shortener, error)
	RedisGet(c *gin.Context, shortened string) (string, error)
	FindAll(sort string) ([]entity.Shortener, error)
}

type shortenerRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewShortener(db *gorm.DB, redis *redis.Client) *shortenerRepository {
	return &shortenerRepository{db, redis}
}

func (r *shortenerRepository) Save(shortener entity.Shortener) (entity.Shortener, error) {
	err := r.db.Create(&shortener).Error
	if err != nil {
		return shortener, err
	}

	return shortener, nil
}

func (r *shortenerRepository) RedisSave(c *gin.Context, shortener entity.Shortener) (entity.Shortener, error) {
	err := r.redis.Set(c, shortener.ID, shortener.TargetURL, 24*time.Hour).Err()
	if err != nil {
		log.Println("Failed to cache URL mapping:", err)
	}

	return shortener, nil
}

func (r *shortenerRepository) RedisGet(c *gin.Context, shortened string) (string, error) {
	target, err := r.redis.Get(c, shortened).Result()
	if err != nil {
		log.Println("Failed to cache URL mapping:", err)
	}

	return target, nil
}

func (r *shortenerRepository) FindByID(url string) (entity.Shortener, error) {
	var shortener entity.Shortener

	err := r.db.Where("id = ?", url).
		First(&shortener).Error
	if err != nil {
		return shortener, err
	}

	shortener.Clicks += 1
	r.db.Save(&shortener)

	return shortener, nil
}

func (r *shortenerRepository) FindAll(sort string) ([]entity.Shortener, error) {
	var shortener []entity.Shortener

	err := r.db.Order("clicks" + " " + sort).
		Find(&shortener).Error

	if err != nil {
		return shortener, err
	}

	return shortener, nil
}
