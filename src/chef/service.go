package chef

import (
	"reku-technical-test/src/entity"
	"reku-technical-test/src/repository"
)

type Service interface {
	StoreChef(request CheftRequest) (entity.Chef, error)
}

type service struct {
	repository repository.ChefRepository
}

func NewService(repository repository.ChefRepository) *service {
	return &service{repository}
}

func (s *service) StoreChef(request CheftRequest) (entity.Chef, error) {
	var chef entity.Chef
	chef.Name = request.Name
	chef.Busy = 0

	newChef, err := s.repository.Save(chef)
	if err != nil {
		return newChef, err
	}

	return newChef, nil
}
