package tenant

import (
	"testing"
	"time"

	"sudhagar/glad/entity"

	"github.com/stretchr/testify/assert"
)

const (
	tenantDefault entity.ID = 13790492210917015554
)

func newFixtureTenant() *entity.Tenant {
	return &entity.Tenant{
		ID:        tenantDefault,
		Username:  "alice@wonder.land",
		Password:  "testing123",
		AuthToken: "token123",
		CreatedAt: time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	tenant := newFixtureTenant()
	_, err := m.CreateTenant(tenant.Username, tenant.Password)
	assert.Nil(t, err)
	assert.False(t, tenant.CreatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	t1 := newFixtureTenant()
	t2 := newFixtureTenant()
	t2.Username = "bob@wunder.land"

	tID, _ := m.CreateTenant(t1.Username, t1.Password)
	_, _ = m.CreateTenant(t2.Username, t2.Password)

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListTenants()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetTenant(tID)
		assert.Nil(t, err)
		// do not compare ID, because it's generated in CreateTenant
		assert.Equal(t, t1.Username, saved.Username)
		// do not compare password, because password is encrypted
		// in CreateTenant
	})
}

func Test_Update(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	tenant := newFixtureTenant()
	id, err := m.CreateTenant(tenant.Username, tenant.Password)
	assert.Nil(t, err)

	saved, _ := m.GetTenant(id)
	saved.Password = "testing456"
	assert.Nil(t, m.UpdateTenant(saved))

	updated, err := m.GetTenant(id)
	assert.Nil(t, err)
	assert.Equal(t, saved.Password, updated.Password)
}

func TestDelete(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)

	t1 := newFixtureTenant()
	t2 := newFixtureTenant()
	t2ID, _ := m.CreateTenant(t2.Username, t2.Password)

	err := m.DeleteTenant(t1.ID)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeleteTenant(t2ID)
	assert.Nil(t, err)
	_, err = m.GetTenant(t2ID)
	assert.Equal(t, entity.ErrNotFound, err)
}

func TestLogin(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)

	// pre-requisities
	t1 := newFixtureTenant()
	t2 := newFixtureTenant()
	t2.Username = "bob@wunder.land"

	t1ID, _ := m.CreateTenant(t1.Username, t1.Password)
	_, _ = m.CreateTenant(t2.Username, t2.Password)

	// test
	t.Run("valid credentials", func(t *testing.T) {
		tenant, err := m.Login(t1.Username, t1.Password)
		assert.Nil(t, err)
		assert.Equal(t, t1.Username, tenant.Username)
		assert.Equal(t, t1ID, tenant.ID)
	})

	t.Run("invalid credentials", func(t *testing.T) {
		tenant, err := m.Login(t1.Username, "PasswordInvalid")
		assert.Equal(t, err, entity.ErrAuthFailure)
		assert.Nil(t, tenant)
	})
}
