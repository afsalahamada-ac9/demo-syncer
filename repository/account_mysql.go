package repository

import (
	"database/sql"
	"time"

	"sudhagar/glad/entity"
)

// AccountMySQL mysql repo
type AccountMySQL struct {
	db *sql.DB
}

// NewAccountMySQL create new repository
func NewAccountMySQL(db *sql.DB) *AccountMySQL {
	return &AccountMySQL{
		db: db,
	}
}

// Create a account
func (r *AccountMySQL) Create(e *entity.Account) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO account (id, tenant_id, username, type, created_at) 
		VALUES(?,?,?,?,?)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		e.ID,
		e.TenantID,
		e.Username,
		int(e.Type),
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

// TODO: GetByID()
// TODO: Need a filter by TenantID as well
// Get a account
func (r *AccountMySQL) Get(username string) (*entity.Account, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, tenant_id, type, created_at FROM account WHERE username = ?`)
	if err != nil {
		return nil, err
	}
	var t entity.Account
	err = stmt.QueryRow(username).Scan(&t.ID, &t.TenantID, &t.Type, &t.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	t.Username = username
	return &t, nil
}

// TODO: UpdateByID()
// Update a account
func (r *AccountMySQL) Update(e *entity.Account) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec(`UPDATE account SET username = ?, type = ?, updated_at = ? WHERE id = ?`,
		e.Username, int(e.Type), e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

// List accounts
func (r *AccountMySQL) List(tenantID entity.ID) ([]*entity.Account, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, tenant_id, username, type, created_at FROM account WHERE tenant_id = ?`)
	if err != nil {
		return nil, err
	}
	var accounts []*entity.Account
	rows, err := stmt.Query(tenantID)
	if err != nil {
		return nil, err
	}

	var username sql.NullString
	for rows.Next() {
		var t entity.Account
		err = rows.Scan(&t.ID, &t.TenantID, &username, &t.Type, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		t.Username = username.String
		accounts = append(accounts, &t)
	}
	return accounts, nil
}

// TODO: DeleteByID()
// Delete a account
func (r *AccountMySQL) Delete(username string) error {
	res, err := r.db.Exec(`DELETE FROM account WHERE username = ?`, username)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Get total accounts
func (r *AccountMySQL) GetCount(tenantID entity.ID) (int, error) {
	stmt, err := r.db.Prepare(`SELECT count(*) FROM account WHERE tenant_id = ?`)
	if err != nil {
		return 0, err
	}

	var count int
	err = stmt.QueryRow(tenantID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
