package services

import (
	"errors"
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductService struct {
	repo    *repositories.ProductRepository
	catRepo *repositories.CategoryRepository
}

func NewProductService(
	repo *repositories.ProductRepository,
	catRepo *repositories.CategoryRepository,
) *ProductService {
	return &ProductService{repo, catRepo}
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Create(product *models.Product) error {
	if product.Name == "" {
		return errors.New("Produk name is required")
	}
	// cek category exist
	_, err := s.catRepo.GetByID(product.CategoryID)
	if err != nil {
		return errors.New("category not found")
	}

	return s.repo.Create(product)
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(product *models.Product) error {
	// cek category exist
	_, err := s.catRepo.GetByID(product.CategoryID)
	if err != nil {
		return errors.New("category not found")
	}

	return s.repo.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
