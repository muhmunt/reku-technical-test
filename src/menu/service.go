package menu

import (
	"reku-technical-test/src/entity"
	"reku-technical-test/src/repository"
)

type Service interface {
	GetAllMenu() ([]entity.Menu, error)
	GetMenuById(id int) (entity.Menu, error)
}

type service struct {
	repository repository.MenuRepository
}

func NewService(repository repository.MenuRepository) *service {
	return &service{repository}
}

func (s *service) GetAllMenu() ([]entity.Menu, error) {
	menu, err := s.repository.FindAll()

	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (s *service) GetMenuById(id int) (entity.Menu, error) {
	menu, err := s.repository.FindByID(id)

	if err != nil {
		return menu, err
	}

	return menu, nil
}
