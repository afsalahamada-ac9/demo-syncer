/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package course

import (
	"strings"
	"time"

	"sudhagar/glad/entity"
)

// Service course usecase
type Service struct {
	repo Repository
}

// NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateCourse creates a course
func (s *Service) CreateCourse(tenantID entity.ID,
	extID string,
	centerID entity.ID,
	productID entity.ID,
	name, notes, timezone string,
	address entity.CourseAddress,
	status entity.CourseStatus,
	mode entity.CourseMode,
	maxAttendees, numAttendees int32,
	isAutoApprove bool,
) (entity.ID, error) {
	c, err := entity.NewCourse(tenantID, extID, centerID, productID,
		name, notes, timezone,
		address, status, mode,
		maxAttendees, numAttendees, isAutoApprove)
	if err != nil {
		return entity.IDInvalid, err
	}
	return s.repo.Create(c)
}

// GetCourse retrieves a course
func (s *Service) GetCourse(id entity.ID) (*entity.Course, error) {
	t, err := s.repo.Get(id)
	if t == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return t, nil
}

// SearchCourses search course
func (s *Service) SearchCourses(tenantID entity.ID,
	query string, page, limit int,
) ([]*entity.Course, error) {
	courses, err := s.repo.Search(tenantID, strings.ToLower(query), page, limit)
	if err != nil {
		return nil, err
	}
	if len(courses) == 0 {
		return nil, entity.ErrNotFound
	}
	return courses, nil
}

// ListCourses list course
func (s *Service) ListCourses(tenantID entity.ID, page, limit int) ([]*entity.Course, error) {
	courses, err := s.repo.List(tenantID, page, limit)
	if err != nil {
		return nil, err
	}
	if len(courses) == 0 {
		return nil, entity.ErrNotFound
	}
	return courses, nil
}

// DeleteCourse Delete a course
func (s *Service) DeleteCourse(id entity.ID) error {
	t, err := s.GetCourse(id)
	if t == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// UpdateCourse Update a course
func (s *Service) UpdateCourse(c *entity.Course) error {
	err := c.Validate()
	if err != nil {
		return err
	}
	c.UpdatedAt = time.Now()
	return s.repo.Update(c)
}

// GetCount gets total course count
func (s *Service) GetCount(tenantID entity.ID) int {
	count, err := s.repo.GetCount(tenantID)
	if err != nil {
		return 0
	}

	return count
}
