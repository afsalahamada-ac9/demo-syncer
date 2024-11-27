/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package course

import (
	"sudhagar/glad/entity"
)

// Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Course, error)
	Search(tenantID entity.ID, query string, page, limit int) ([]*entity.Course, error)
	List(tenantID entity.ID, page, limit int) ([]*entity.Course, error)
	GetCount(id entity.ID) (int, error)
}

// Writer course writer
type Writer interface {
	Create(e *entity.Course) (entity.ID, error)
	Update(e *entity.Course) error
	Delete(id entity.ID) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	GetCourse(id entity.ID) (*entity.Course, error)
	SearchCourses(tenantID entity.ID, query string, page, limit int) ([]*entity.Course, error)
	ListCourses(tenantID entity.ID, page, limit int) ([]*entity.Course, error)
	CreateCourse(tenantID entity.ID,
		extID string,
		centerID entity.ID,
		name, notes, timezone string,
		address entity.CourseAddress,
		status entity.CourseStatus,
		mode entity.CourseMode,
		maxAttendees, numAttendees int32,
		isAutoApprove bool,
	) (entity.ID, error)
	UpdateCourse(e *entity.Course) error
	DeleteCourse(id entity.ID) error
	GetCount(id entity.ID) int
}
