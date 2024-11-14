/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package account

import (
	"sudhagar/glad/entity"
)

// Reader interface
type Reader interface {
	Get(username string) (*entity.Account, error)
	List(tenantID entity.ID) ([]*entity.Account, error)
	GetCount(tenantId entity.ID) (int, error)
}

// Writer interface
type Writer interface {
	Create(e *entity.Account) error
	Update(e *entity.Account) error
	Delete(username string) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// Messager interface
type Messager interface {
	// Start starts the messager interface; In future, type could
	// be passed, so appropriate service such as WA, Msgr, could be
	// started
	Start() (username string, qrData string, err error) // data is optional
	GetStatus(username string) (entity.AccountStatus, error)
	Stop(username string) error // error may not be too relevant here
}

// UseCase interface
type UseCase interface {
	GetQR(
		tenantID entity.ID,
		at entity.AccountType) (username string, qrData string, err error)
	GetStatus(username string, tenantID entity.ID) (entity.AccountStatus, error)

	CreateAccount(
		tenantID entity.ID,
		username string,
		at entity.AccountType) error
	GetAccount(username string) (*entity.Account, error)
	ListAccounts(tenantID entity.ID) ([]*entity.Account, error)
	UpdateAccount(e *entity.Account) error
	DeleteAccount(username string) error
	GetCount(tenantId entity.ID) int
}
