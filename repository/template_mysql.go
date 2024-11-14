package repository

import (
	"database/sql"
	"time"

	"sudhagar/glad/entity"
)

// TemplateMySQL mysql repo
type TemplateMySQL struct {
	db *sql.DB
}

// NewTemplateMySQL create new repository
func NewTemplateMySQL(db *sql.DB) *TemplateMySQL {
	return &TemplateMySQL{
		db: db,
	}
}

// Create a template
func (r *TemplateMySQL) Create(e *entity.Template) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO template (id, tenant_id, name, type, content, created_at) 
		VALUES(?,?,?,?,?,?)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.TenantID,
		e.Name,
		int(e.Type),
		e.Content,
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

// Get a template
func (r *TemplateMySQL) Get(id entity.ID) (*entity.Template, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, tenant_id, name, type, content, created_at FROM template WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	var t entity.Template
	var name sql.NullString
	err = stmt.QueryRow(id).Scan(&t.ID, &t.TenantID, &name, &t.Type, &t.Content, &t.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	t.Name = name.String
	return &t, nil
}

// Update a template
func (r *TemplateMySQL) Update(e *entity.Template) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec(`
		UPDATE template SET name = ?, content = ?, type = ?, updated_at = ? WHERE id = ?`,
		e.Name, e.Content, int(e.Type), e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

// Search templates
func (r *TemplateMySQL) Search(tenantID entity.ID,
	query string,
) ([]*entity.Template, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, tenant_id, name, type, content, created_at FROM template
		WHERE tenant_id = ? AND content LIKE ?`)
	if err != nil {
		return nil, err
	}
	var templates []*entity.Template
	rows, err := stmt.Query(tenantID, "%"+query+"%")
	if err != nil {
		return nil, err
	}

	var name sql.NullString
	for rows.Next() {
		var t entity.Template
		err = rows.Scan(&t.ID, &t.TenantID, &name, &t.Type, &t.Content, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		t.Name = name.String
		templates = append(templates, &t)
	}

	return templates, nil
}

// List templates
func (r *TemplateMySQL) List(tenantID entity.ID) ([]*entity.Template, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, tenant_id, name, type, content, created_at FROM template WHERE tenant_id = ?`)
	if err != nil {
		return nil, err
	}
	var templates []*entity.Template
	rows, err := stmt.Query(tenantID)
	if err != nil {
		return nil, err
	}

	var name sql.NullString
	for rows.Next() {
		var t entity.Template
		err = rows.Scan(&t.ID, &t.TenantID, &name, &t.Type, &t.Content, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		t.Name = name.String
		templates = append(templates, &t)
	}
	return templates, nil
}

// Delete a template
func (r *TemplateMySQL) Delete(id entity.ID) error {
	res, err := r.db.Exec(`DELETE FROM template WHERE id = ?`, id)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Get total templates
func (r *TemplateMySQL) GetCount(tenantID entity.ID) (int, error) {
	stmt, err := r.db.Prepare(`SELECT count(*) FROM template WHERE tenant_id = ?`)
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
