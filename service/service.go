package service

import (
	"customer/models"
	"customer/repository"
)

type CustomerService struct {
	Repo *repository.CustomerRepository
}

func (s *CustomerService) GetAll() ([]models.Customer, error) {
	return s.Repo.GetAll()
}

func (s *CustomerService) GetByID(id int) (*models.Customer, error) {
	return s.Repo.GetByID(id)
}

func (s *CustomerService) Create(c *models.Customer) error {
	return s.Repo.Create(c)
}

func (s *CustomerService) Update(id int, c *models.Customer) error {
	return s.Repo.Update(id, c)
}

func (s *CustomerService) Delete(id int) error {
	return s.Repo.Delete(id)
}
