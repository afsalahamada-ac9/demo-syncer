package account

import (
	"testing"
	"time"

	"sudhagar/glad/entity"

	"github.com/stretchr/testify/assert"
)

const (
	accountUsernameAlice  string = "12345550001"
	accountUsername2Alice string = "12345550002"

	accountIDAlice  entity.ID = 13790492210917010000
	accountID2Alice entity.ID = 13790492210917010002
	tenantAlice     entity.ID = 13790492210917015554

	// todo: add multi-tenant support
	// tenantBob entity.ID = 13790492210917015555
)

func newFixtureAccount() *entity.Account {
	return &entity.Account{
		ID:        accountIDAlice,
		TenantID:  tenantAlice,
		Username:  accountUsernameAlice,
		Type:      entity.AccountWhatsApp,
		CreatedAt: time.Now(),
	}
}

// Mock messager
type messager struct {
}

func (m *messager) Start() (username string, qrData string, err error) {
	return accountUsernameAlice, "", nil
}

func (m *messager) Stop(id string) error {
	return nil
}

func (m *messager) GetStatus(username string) (entity.AccountStatus, error) {
	return entity.AccountStatusUnknown, nil
}

// End of Mock messager

func Test_Create(t *testing.T) {
	repo := newInmem()
	m := NewService(repo, &messager{})
	account := newFixtureAccount()
	err := m.CreateAccount(account.TenantID, account.Username, account.Type)
	assert.Nil(t, err)
	assert.False(t, account.CreatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := newInmem()
	m := NewService(repo, &messager{})
	account1 := newFixtureAccount()
	account2 := newFixtureAccount()
	account2.ID = accountID2Alice
	account2.Username = accountUsername2Alice

	_ = m.CreateAccount(account1.TenantID,
		account1.Username,
		account1.Type)
	_ = m.CreateAccount(account2.TenantID,
		account2.Username,
		account2.Type)

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListAccounts(account1.TenantID)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetAccount(account1.Username)
		assert.Nil(t, err)
		assert.Equal(t, account1.TenantID, saved.TenantID)
		assert.Equal(t, account1.Type, saved.Type)
		assert.Equal(t, account1.Username, saved.Username)
	})
}

// It's unlikely that the update will be called in this entity model.
// Perhaps a human readable name can be given for customer to reference.
func Test_Update(t *testing.T) {
	repo := newInmem()
	m := NewService(repo, &messager{})
	account := newFixtureAccount()
	err := m.CreateAccount(
		account.TenantID,
		account.Username,
		account.Type)
	assert.Nil(t, err)

	saved, _ := m.GetAccount(account.Username)
	saved.Username = "starred"
	assert.Nil(t, m.UpdateAccount(saved))

	_, err = m.GetAccount(account.Username)
	assert.Equal(t, entity.ErrNotFound, err)

	updated, err := m.GetAccount(saved.Username)
	assert.Nil(t, err)
	assert.Equal(t, "starred", updated.Username)
}

func TestDelete(t *testing.T) {
	repo := newInmem()
	m := NewService(repo, &messager{})

	account1 := newFixtureAccount()

	account2 := newFixtureAccount()
	account2.ID = accountID2Alice
	account2.Username = accountUsername2Alice
	_ = m.CreateAccount(
		account2.TenantID,
		account2.Username,
		account2.Type)

	err := m.DeleteAccount(account1.Username)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeleteAccount(account2.Username)
	assert.Nil(t, err)

	_, err = m.GetAccount(account2.Username)
	assert.Equal(t, entity.ErrNotFound, err)
}
