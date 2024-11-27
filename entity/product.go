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
	ID               ID                `json:"id"`
	ExtID            string            `json:"extId"`
	TenantID         ID                `json:"tenantId"`
	ExtName          string            `json:"extName"`
	Title            string            `json:"title"`
	CType            string            `json:"ctype"`
	BaseProductExtID string            `json:"baseProductExtId,omitempty"`
	DurationDays     int32             `json:"durationDays,omitempty"`
	Visibility       ProductVisibility `json:"visibility,omitempty"`
	MaxAttendees     int32             `json:"maxAttendees,omitempty"`
	Format           ProductFormat     `json:"format,omitempty"`
	CreatedAt        time.Time         `json:"createdAt"`
	UpdatedAt        time.Time         `json:"updatedAt"`
}

// NewProduct creates a new product with the given parameters
func NewProduct(tenantID ID,
	extID string,
	extName string,
	title string,
	ctype string,
	baseProductExtID string,
	durationDays int32,
	visibility ProductVisibility,
	maxAttendees int32,
	format ProductFormat,
) (*Product, error) {
	p := &Product{
		ID:               NewID(),
		ExtID:            extID,
		TenantID:         tenantID,
		ExtName:          extName,
		Title:            title,
		CType:            ctype,
		BaseProductExtID: baseProductExtID,
		DurationDays:     durationDays,
		Visibility:       visibility,
		MaxAttendees:     maxAttendees,
		Format:           format,
		CreatedAt:        time.Now(),
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

	if p.ExtName == "" {
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
