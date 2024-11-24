package product

import (
	"strings"
	"time"

	"sudhagar/glad/entity"
)

// Service product usecase
type Service struct {
	repo Repository
}

// NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateProduct creates a product
func (s *Service) CreateProduct(tenantID entity.ID,
	extID string,
	name string,
	title string,
	ctype string,
	baseProductID string,
	durationDays int32,
	visibility entity.ProductVisibility,
	maxAttendees int32,
	format entity.ProductFormat,
	isDeleted bool,
) (entity.ID, error) {
	p, err := entity.NewProduct(tenantID,
		extID,
		name,
		title,
		ctype,
		baseProductID,
		durationDays,
		visibility,
		maxAttendees,
		format,
		isDeleted)
	if err != nil {
		return entity.IDInvalid, err
	}

	return s.repo.Create(p)
}

// GetProduct retrieves a product
func (s *Service) GetProduct(id entity.ID) (*entity.Product, error) {
	p, err := s.repo.Get(id)
	if p == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return p, nil
}

// SearchProducts search product
func (s *Service) SearchProducts(tenantID entity.ID, query string) ([]*entity.Product, error) {
	products, err := s.repo.Search(tenantID, strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, entity.ErrNotFound
	}
	return products, nil
}

// ListProducts list products
func (s *Service) ListProducts(tenantID entity.ID) ([]*entity.Product, error) {
	products, err := s.repo.List(tenantID, 0, 0)
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, entity.ErrNotFound
	}
	return products, nil
}

// UpdateProduct Update a product
func (s *Service) UpdateProduct(p *entity.Product) error {
	err := p.Validate()
	if err != nil {
		return err
	}
	p.UpdatedAt = time.Now()
	return s.repo.Update(p)
}

// DeleteProduct Delete a product
func (s *Service) DeleteProduct(id entity.ID) error {
	p, err := s.GetProduct(id)
	if p == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// GetCount gets total product count
func (s *Service) GetCount(tenantID entity.ID) int {
	count, err := s.repo.GetCount(tenantID)
	if err != nil {
		return 0
	}

	return count
}
