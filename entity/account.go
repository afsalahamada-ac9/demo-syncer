/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

import (
	"time"
)

// Account type
type AccountType int

const (
	AccountWhatsApp AccountType = iota
	// Add new types here
)

// AccountStatus type
type AccountStatus int

const (
	AccountStatusUnknown AccountStatus = iota
	AccountStatusConnected
	AccountStatusLoggedIn
	// Add new statuses here
)

// Account data
type Account struct {
	ID       ID
	TenantID ID
	Username string
	Type     AccountType

	// TODO: Add registration data here

	// meta data
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewAccount create a new account
func NewAccount(tenantID ID, username string, at AccountType) (*Account, error) {
	t := &Account{
		ID:        NewID(),
		TenantID:  tenantID,
		Username:  username,
		Type:      at,
		CreatedAt: time.Now(),
	}
	err := t.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return t, nil
}

// Validate validate account
func (t *Account) Validate() error {
	if t.Username == "" {
		return ErrInvalidEntity
	}

	return nil
}
