/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package repository

import (
	"database/sql"
	"time"

	"sudhagar/glad/entity"
)

// TenantMySQL mysql repo
type TenantMySQL struct {
	db *sql.DB
}

// NewTenantMySQL create new repository
func NewTenantMySQL(db *sql.DB) *TenantMySQL {
	return &TenantMySQL{
		db: db,
	}
}

// Create a Tenant
func (r *TenantMySQL) Create(e *entity.Tenant) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO tenant (id, username, password, created_at) 
		VALUES(?,?,?,?)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Username,
		e.Password,
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return e.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

// Get a Tenant
func (r *TenantMySQL) Get(id entity.ID) (*entity.Tenant, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, username, password, auth_token, created_at FROM tenant WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	var t entity.Tenant
	var token sql.NullString

	err = stmt.QueryRow(id).Scan(&t.ID, &t.Username, &t.Password, &token, &t.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	t.AuthToken = token.String
	return &t, nil
}

// Get a Tenant by username
func (r *TenantMySQL) GetByName(username string) (*entity.Tenant, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, username, password, auth_token, created_at FROM tenant WHERE username = ?`)
	if err != nil {
		return nil, err
	}
	var t entity.Tenant
	var token sql.NullString

	err = stmt.QueryRow(username).Scan(&t.ID, &t.Username, &t.Password, &token, &t.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	t.AuthToken = token.String
	return &t, nil
}

// Update a Tenant
func (r *TenantMySQL) Update(e *entity.Tenant) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec(`UPDATE tenant SET username = ?, auth_token = ?, updated_at = ? WHERE id = ?`,
		e.Username, e.AuthToken, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

// List Tenants
func (r *TenantMySQL) List() ([]*entity.Tenant, error) {
	stmt, err := r.db.Prepare(`SELECT id, username, auth_token, created_at FROM tenant`)
	if err != nil {
		return nil, err
	}
	var Tenants []*entity.Tenant
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t entity.Tenant
		var token sql.NullString
		err = rows.Scan(&t.ID, &t.Username, &token, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		t.AuthToken = token.String
		Tenants = append(Tenants, &t)
	}
	return Tenants, nil
}

// Delete a Tenant
func (r *TenantMySQL) Delete(id entity.ID) error {
	res, err := r.db.Exec(`DELETE FROM tenant WHERE id = ?`, id)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Get total tenants
func (r *TenantMySQL) GetCount() (int, error) {
	var count int
	err := r.db.QueryRow(`SELECT count(*) FROM tenant`).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
