/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

import (
	"time"
)

// Template type
type TemplateType int

const (
	TemplateText TemplateType = iota
	// Add new types here
)

// Template data
type Template struct {
	ID       ID
	TenantID ID
	Name     string
	Type     TemplateType
	Content  string

	// meta data
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewTemplate create a new template
func NewTemplate(tenantID ID, name string, tt TemplateType, content string) (*Template, error) {
	t := &Template{
		ID:        NewID(),
		TenantID:  tenantID,
		Name:      name,
		Type:      tt,
		Content:   content,
		CreatedAt: time.Now(),
	}
	err := t.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return t, nil
}

// Validate validate template
func (t *Template) Validate() error {
	if t.Content == "" {
		return ErrInvalidEntity
	}
	return nil
}
