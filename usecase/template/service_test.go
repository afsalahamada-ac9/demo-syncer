package template

import (
	"testing"
	"time"

	"sudhagar/glad/entity"

	"github.com/stretchr/testify/assert"
)

const (
	templateDefault entity.ID = 13790493495087071234
	tenantAlice     entity.ID = 13790492210917015554

	// todo: add multi-tenant support
	// tenantBob    entity.ID = 13790492210917015555
)

func newFixtureTemplate() *entity.Template {
	return &entity.Template{
		ID:        templateDefault,
		TenantID:  tenantAlice,
		Name:      "default1",
		Type:      entity.TemplateText,
		Content:   "This is a default test message",
		CreatedAt: time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	tmpl := newFixtureTemplate()
	_, err := m.CreateTemplate(tmpl.TenantID, tmpl.Name, tmpl.Type, tmpl.Content)
	assert.Nil(t, err)
	assert.False(t, tmpl.CreatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	tmpl1 := newFixtureTemplate()
	tmpl2 := newFixtureTemplate()
	tmpl2.Name = "default2"
	tmpl2.Content = "This is a test message #2"

	tID, _ := m.CreateTemplate(tmpl1.TenantID, tmpl1.Name, tmpl1.Type, tmpl1.Content)
	_, _ = m.CreateTemplate(tmpl2.TenantID, tmpl1.Name, tmpl2.Type, tmpl2.Content)

	t.Run("search", func(t *testing.T) {
		res, err := m.SearchTemplates(tmpl1.TenantID, "default")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, tmpl1.Content, res[0].Content)

		res, err = m.SearchTemplates(tmpl1.TenantID, "non-existent")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, res)
	})
	t.Run("list all", func(t *testing.T) {
		all, err := m.ListTemplates(tmpl1.TenantID)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetTemplate(tID)
		assert.Nil(t, err)
		assert.Equal(t, tmpl1.TenantID, saved.TenantID)
		assert.Equal(t, tmpl1.Type, saved.Type)
		assert.Equal(t, tmpl1.Content, saved.Content)
	})
}

func Test_Update(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	tmpl := newFixtureTemplate()
	id, err := m.CreateTemplate(tmpl.TenantID, tmpl.Name, tmpl.Type, tmpl.Content)
	assert.Nil(t, err)

	saved, _ := m.GetTemplate(id)
	saved.Content = "This is an updated message"
	assert.Nil(t, m.UpdateTemplate(saved))

	updated, err := m.GetTemplate(id)
	assert.Nil(t, err)
	assert.Equal(t, "This is an updated message", updated.Content)
}

func TestDelete(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)

	tmpl1 := newFixtureTemplate()
	tmpl2 := newFixtureTemplate()
	t2ID, _ := m.CreateTemplate(tmpl2.TenantID, tmpl2.Name, tmpl2.Type, tmpl2.Content)

	err := m.DeleteTemplate(tmpl1.ID)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeleteTemplate(t2ID)
	assert.Nil(t, err)
	_, err = m.GetTemplate(t2ID)
	assert.Equal(t, entity.ErrNotFound, err)
}
