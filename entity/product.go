/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

import "time"

// ProductVisibility represents the visibility status of a product
type ProductVisibility string

const (
	ProductVisibilityPublic   ProductVisibility = "Public"
	ProductVisibilityUnlisted ProductVisibility = "Unlisted"
)

// ProductFormat represents the format of a product
type ProductFormat string

const (
	ProductFormatInPerson    ProductFormat = "In Person"
	ProductFormatOnline      ProductFormat = "Online"
	ProductFormatDestination ProductFormat = "Destination Retreats"
)

// Product represents a product entity
// TODO: json tags must be moved to presenter
type Product struct {
	ID            ID                `json:"id"`
	ExtID         string            `json:"extId"`
	TenantID      ID                `json:"tenantId"`
	Name          string            `json:"name"`
	Title         string            `json:"title"`
	CType         string            `json:"ctype"`
	BaseProductID string            `json:"baseProductId,omitempty"`
	DurationDays  int32             `json:"durationDays,omitempty"`
	Visibility    ProductVisibility `json:"visibility,omitempty"`
	MaxAttendees  int32             `json:"maxAttendees,omitempty"`
	Format        ProductFormat     `json:"format,omitempty"`
	IsDeleted     bool              `json:"isDeleted"`
	CreatedAt     time.Time         `json:"createdAt"`
	UpdatedAt     time.Time         `json:"updatedAt"`
}

// NewProduct creates a new product with the given parameters
func NewProduct(tenantID ID,
	extID string,
	name string,
	title string,
	ctype string,
	baseProductID string,
	durationDays int32,
	visibility ProductVisibility,
	maxAttendees int32,
	format ProductFormat,
	isDeleted bool,
) (*Product, error) {
	p := &Product{
		ID:            NewID(),
		ExtID:         extID,
		TenantID:      tenantID,
		Name:          name,
		Title:         title,
		CType:         ctype,
		BaseProductID: baseProductID,
		DurationDays:  durationDays,
		Visibility:    visibility,
		MaxAttendees:  maxAttendees,
		Format:        format,
		IsDeleted:     false,
		CreatedAt:     time.Now(),
	}

	err := p.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return p, nil
}

// Validate validates the product fields
func (p *Product) Validate() error {
	if p.TenantID == 0 {
		return ErrInvalidEntity
	}

	if p.Name == "" {
		return ErrInvalidEntity
	}

	if p.Title == "" {
		return ErrInvalidEntity
	}

	if p.CType == "" {
		return ErrInvalidEntity
	}

	return nil
}
