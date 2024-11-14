package account

import (
	"time"

	"sudhagar/glad/entity"
)

// Service account usecase
type Service struct {
	repo Repository
	msgr Messager
}

// NewService create new service
func NewService(r Repository, m Messager) *Service {
	return &Service{
		repo: r,
		msgr: m,
	}
}

// GetQR gets QR data to add an account
func (s *Service) GetQR(
	tenantID entity.ID,
	at entity.AccountType) (username string, qrData string, err error) {

	// TODO: Validate account type
	// TODO: Store tenant id?

	username, qrData, err = s.msgr.Start()

	return username, qrData, err
}

func (s *Service) GetStatus(username string, tenantID entity.ID) (entity.AccountStatus, error) {
	// TODO: Tenant id to be implemented
	return s.msgr.GetStatus(username)
}

// CreateAccount create a account
func (s *Service) CreateAccount(
	tenantID entity.ID,
	username string,
	at entity.AccountType) error {
	account, err := entity.NewAccount(tenantID, username, at)
	if err != nil {
		return err
	}
	return s.repo.Create(account)
}

// GetAccount get a account
func (s *Service) GetAccount(username string) (*entity.Account, error) {
	account, err := s.repo.Get(username)
	if account == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return account, nil
}

// ListAccounts list account
func (s *Service) ListAccounts(tenantID entity.ID) ([]*entity.Account, error) {
	accounts, err := s.repo.List(tenantID)
	if err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		return nil, entity.ErrNotFound
	}
	return accounts, nil
}

// UpdateAccount Update a account
func (s *Service) UpdateAccount(t *entity.Account) error {
	err := t.Validate()
	if err != nil {
		return err
	}
	t.UpdatedAt = time.Now()
	return s.repo.Update(t)
}

// DeleteAccount Delete a account
func (s *Service) DeleteAccount(username string) error {
	account, err := s.GetAccount(username)
	if account == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}

	return s.repo.Delete(username)
}

// GetCount gets total account count
func (s *Service) GetCount(tenantID entity.ID) int {
	count, err := s.repo.GetCount(tenantID)
	if err != nil {
		return 0
	}

	return count
}
