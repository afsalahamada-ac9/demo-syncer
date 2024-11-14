/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package template

import (
	"sudhagar/glad/entity"
)

// Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Template, error)
	Search(tenantID entity.ID, query string) ([]*entity.Template, error)
	List(tenantID entity.ID) ([]*entity.Template, error)
	GetCount(id entity.ID) (int, error)
}

// Writer template writer
type Writer interface {
	Create(e *entity.Template) (entity.ID, error)
	Update(e *entity.Template) error
	Delete(id entity.ID) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	GetTemplate(id entity.ID) (*entity.Template, error)
	SearchTemplates(tenantID entity.ID, query string) ([]*entity.Template, error)
	ListTemplates(tenantID entity.ID) ([]*entity.Template, error)
	CreateTemplate(tenantID entity.ID, name string, tt entity.TemplateType, content string) (entity.ID, error)
	UpdateTemplate(e *entity.Template) error
	DeleteTemplate(id entity.ID) error
	GetCount(id entity.ID) int
}
