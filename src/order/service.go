package order

import (
	"errors"
	"fmt"
	"reku-technical-test/src/entity"
	"reku-technical-test/src/helper"
	"reku-technical-test/src/repository"
	"sync"
	"time"
)

var (
	mu sync.Mutex
)

type Service interface {
	PlaceOrder(request entity.Menu) (entity.Order, error)
}

type service struct {
	repository     repository.OrderRepository
	chefRepository repository.ChefRepository
}

func NewService(repository repository.OrderRepository, chefRepository repository.ChefRepository) *service {
	return &service{repository, chefRepository}
}

func (s *service) PlaceOrder(request entity.Menu) (entity.Order, error) {
	var order entity.Order

	order.Duration = request.Duration
	order.Pizza = request.Name
	order.Status = "processing"
	order.CreatedAt = helper.TimeNow()

	newOrder, err := s.repository.Save(order)

	if err != nil {
		return newOrder, err
	}

	chef, err := s.GetAvailableChef()

	if err != nil {
		return entity.Order{}, errors.New("All chefs are busy.")
	}

	_ = s.chefRepository.UpdateStatus(chef.ID, 1)

	go s.ProcessOrder(newOrder, *chef)

	return newOrder, nil
}

func (s *service) GetAvailableChef() (*entity.Chef, error) {
	var availableChef entity.Chef

	availableChef, err := s.chefRepository.FindAvailable()

	if err != nil {
		return nil, err
	}
	return &availableChef, nil
}

func (s *service) ProcessOrder(order entity.Order, chef entity.Chef) {
	mu.Lock()
	defer mu.Unlock()

	time.Sleep(time.Duration(order.Duration) * time.Second)
	fmt.Printf("Order %s is ready\n", order.Pizza)

	order.Status = "ready"

	_ = s.repository.UpdateStatus(order.ID)

	_ = s.chefRepository.UpdateStatus(chef.ID, 0)
}
