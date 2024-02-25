package shortener

import (
	"reku-technical-test/src/entity"
	"reku-technical-test/src/helper"
	"reku-technical-test/src/repository"

	"time"

	"github.com/gin-gonic/gin"
)

type Service interface {
	StoreShortUrl(c *gin.Context, request entity.Shortener) (entity.Shortener, error)
	GetShortUrl(c *gin.Context, request FetchRequest) (string, error)
	GetAllUrl(sort string) ([]entity.Shortener, error)
}

type service struct {
	repository repository.ShortenerRepository
}

func NewService(repository repository.ShortenerRepository) *service {
	return &service{repository}
}

func (s *service) StoreShortUrl(c *gin.Context, shortener entity.Shortener) (entity.Shortener, error) {
	currentTime := time.Now()

	currentTimePlusOneDay := currentTime.Add(24 * time.Hour)

	epochTimeSeconds := currentTimePlusOneDay.Unix()

	shortener.ExpiryDate = int(epochTimeSeconds)

	shortener.CreatedAt = helper.TimeNow()

	newShort, err := s.repository.Save(shortener)
	if err != nil {
		return newShort, err
	}

	_, err = s.repository.RedisSave(c, shortener)
	if err != nil {
		return newShort, err
	}

	return newShort, nil
}

func (s *service) GetShortUrl(c *gin.Context, shortener FetchRequest) (string, error) {
	short, err := s.repository.FindByID(shortener.Shortened)

	if err != nil {
		return "", err
	}

	targetUrl, err := s.repository.RedisGet(c, shortener.Shortened)

	if err == nil {
		return targetUrl, err
	}

	_, err = s.repository.RedisSave(c, short)
	if err != nil {
		return "", err
	}

	return short.TargetURL, nil
}

func (s *service) GetAllUrl(sort string) ([]entity.Shortener, error) {
	short, err := s.repository.FindAll(sort)

	if err != nil {
		return short, err
	}

	return short, nil
}
