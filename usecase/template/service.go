package template

import (
	"strings"
	"time"

	"sudhagar/glad/entity"
)

// Service template usecase
type Service struct {
	repo Repository
}

// NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateTemplate create a template
func (s *Service) CreateTemplate(tenantID entity.ID,
	name string,
	tt entity.TemplateType,
	content string,
) (entity.ID, error) {
	t, err := entity.NewTemplate(tenantID, name, tt, content)
	if err != nil {
		return t.ID, err
	}
	return s.repo.Create(t)
}

// GetTemplate get a template
func (s *Service) GetTemplate(id entity.ID) (*entity.Template, error) {
	t, err := s.repo.Get(id)
	if t == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return t, nil
}

// SearchTemplates search template
func (s *Service) SearchTemplates(tenantID entity.ID,
	query string,
) ([]*entity.Template, error) {
	templates, err := s.repo.Search(tenantID, strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(templates) == 0 {
		return nil, entity.ErrNotFound
	}
	return templates, nil
}

// ListTemplates list template
func (s *Service) ListTemplates(tenantID entity.ID) ([]*entity.Template, error) {
	templates, err := s.repo.List(tenantID)
	if err != nil {
		return nil, err
	}
	if len(templates) == 0 {
		return nil, entity.ErrNotFound
	}
	return templates, nil
}

// DeleteTemplate Delete a template
func (s *Service) DeleteTemplate(id entity.ID) error {
	t, err := s.GetTemplate(id)
	if t == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// UpdateTemplate Update a template
func (s *Service) UpdateTemplate(t *entity.Template) error {
	err := t.Validate()
	if err != nil {
		return err
	}
	t.UpdatedAt = time.Now()
	return s.repo.Update(t)
}

// GetCount gets total template count
func (s *Service) GetCount(tenantID entity.ID) int {
	count, err := s.repo.GetCount(tenantID)
	if err != nil {
		return 0
	}

	return count
}
