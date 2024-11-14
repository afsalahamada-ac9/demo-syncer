/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package template

import (
	"strings"

	"sudhagar/glad/entity"
)

// inmem in memory repo
type inmem struct {
	m map[entity.ID]*entity.Template
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[entity.ID]*entity.Template{}
	return &inmem{
		m: m,
	}
}

// Create a template
func (r *inmem) Create(e *entity.Template) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get a template
func (r *inmem) Get(id entity.ID) (*entity.Template, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

// Update a template
func (r *inmem) Update(e *entity.Template) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// Search templates
func (r *inmem) Search(tenantID entity.ID,
	query string,
) ([]*entity.Template, error) {
	var d []*entity.Template
	for _, j := range r.m {
		if j.TenantID == tenantID &&
			strings.Contains(strings.ToLower(j.Content), query) {
			d = append(d, j)
		}
	}
	return d, nil
}

// List templates
func (r *inmem) List(tenantID entity.ID) ([]*entity.Template, error) {
	var d []*entity.Template
	for _, j := range r.m {
		if j.TenantID == tenantID {
			d = append(d, j)
		}
	}
	return d, nil
}

// Delete a template
func (r *inmem) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	delete(r.m, id)
	return nil
}

// GetCount gets total templates for a given tenant
func (r *inmem) GetCount(tenantID entity.ID) (int, error) {
	count := 0
	for _, j := range r.m {
		if j.TenantID == tenantID {
			count++
		}
	}
	return count, nil
}
