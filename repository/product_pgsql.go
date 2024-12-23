package repository

import (
	"database/sql"
	"time"

	"sudhagar/glad/entity"
)

// ProductPGSQL postgres repo
type ProductPGSQL struct {
	db *sql.DB
}

// NewProductPGSQL create new repository
func NewProductPGSQL(db *sql.DB) *ProductPGSQL {
	return &ProductPGSQL{
		db: db,
	}
}

// Create creates a product
func (r *ProductPGSQL) Create(e *entity.Product) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO product (id, ext_id, tenant_id, ext_name, title, ctype, base_product_ext_id, 
			duration_days, visibility, max_attendees, format, is_auto_approve, created_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.ExtID,
		e.TenantID,
		e.ExtName,
		e.Title,
		e.CType,
		e.BaseProductExtID,
		e.DurationDays,
		string(e.Visibility),
		e.MaxAttendees,
		string(e.Format),
		e.IsAutoApprove,
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

// Get retrieves a product
func (r *ProductPGSQL) Get(id entity.ID) (*entity.Product, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, tenant_id, ext_id, ext_name, title, ctype, base_product_ext_id, 
			duration_days, visibility, max_attendees, format, is_auto_approve, created_at 
		FROM product WHERE id = $1;`)
	if err != nil {
		return nil, err
	}

	var p entity.Product
	var ext_id, base_product_ext_id, visibility, format sql.NullString
	var duration_days, max_attendees sql.NullInt32

	err = stmt.QueryRow(id).Scan(
		&p.ID,
		&p.TenantID,
		&ext_id,
		&p.ExtName,
		&p.Title,
		&p.CType,
		&base_product_ext_id,
		&duration_days,
		&visibility,
		&max_attendees,
		&format,
		&p.IsAutoApprove,
		&p.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	p.ExtID = ext_id.String
	p.BaseProductExtID = base_product_ext_id.String
	p.DurationDays = duration_days.Int32
	p.Visibility = entity.ProductVisibility(visibility.String)
	p.MaxAttendees = max_attendees.Int32
	p.Format = entity.ProductFormat(format.String)

	return &p, nil
}

// Update updates a product
func (r *ProductPGSQL) Update(e *entity.Product) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec(`
		UPDATE product 
		SET ext_name = $1, title = $2, ctype = $3, base_product_ext_id = $4,
			duration_days = $5, visibility = $6, max_attendees = $7,
			format = $8,  is_auto_approve = $9, updated_at = $10
		WHERE id = $11;`,
		e.ExtName,
		e.Title,
		e.CType,
		e.BaseProductExtID,
		e.DurationDays,
		string(e.Visibility),
		e.MaxAttendees,
		string(e.Format),
		e.IsAutoApprove,
		e.UpdatedAt.Format("2006-01-02"),
		e.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

// Search searches products
func (r *ProductPGSQL) Search(tenantID entity.ID, q string, page, limit int) ([]*entity.Product, error) {
	query := `
		SELECT id, tenant_id, ext_id, ext_name, title, ctype, base_product_ext_id,
			duration_days, visibility, max_attendees, format, is_auto_approve, created_at
		FROM product 
		WHERE tenant_id = $1 AND (LOWER(ext_name) LIKE LOWER($2) OR LOWER(title) LIKE LOWER($2))
	`

	// Add pagination if specified
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query += ` LIMIT $3 OFFSET $4;`
		stmt, err := r.db.Prepare(query)
		if err != nil {
			return nil, err
		}
		rows, err := stmt.Query(tenantID, "%"+q+"%", limit, offset)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		return r.scanRows(rows)
	}

	stmt, err := r.db.Prepare(query + ";")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(tenantID, "%"+q+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return r.scanRows(rows)
}

// List lists products
func (r *ProductPGSQL) List(tenantID entity.ID, page, limit int) ([]*entity.Product, error) {
	query := `
		SELECT id, tenant_id, ext_id, ext_name, title, ctype, base_product_ext_id,
			duration_days, visibility, max_attendees, format, is_auto_approve, created_at
		FROM product 
		WHERE tenant_id = $1
	`

	// Add pagination if specified
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query += ` LIMIT $2 OFFSET $3;`
		stmt, err := r.db.Prepare(query)
		if err != nil {
			return nil, err
		}
		rows, err := stmt.Query(tenantID, limit, offset)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		return r.scanRows(rows)
	}

	stmt, err := r.db.Prepare(query + ";")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(tenantID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return r.scanRows(rows)
}

// Delete deletes a product
func (r *ProductPGSQL) Delete(id entity.ID) error {
	res, err := r.db.Exec(`DELETE FROM product WHERE id = $1;`, id)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// GetCount gets total products count for a tenant
func (r *ProductPGSQL) GetCount(tenantID entity.ID) (int, error) {
	stmt, err := r.db.Prepare(`
		SELECT COUNT(*) 
		FROM product 
		WHERE tenant_id = $1;
	`)
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

// scanRows is a helper function to scan rows into product slices
func (r *ProductPGSQL) scanRows(rows *sql.Rows) ([]*entity.Product, error) {
	var products []*entity.Product

	for rows.Next() {
		var p entity.Product
		var ext_id, base_product_ext_id, visibility, format sql.NullString
		var duration_days, max_attendees sql.NullInt32

		err := rows.Scan(
			&p.ID,
			&p.TenantID,
			&ext_id,
			&p.ExtName,
			&p.Title,
			&p.CType,
			&base_product_ext_id,
			&duration_days,
			&visibility,
			&max_attendees,
			&format,
			&p.IsAutoApprove,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		p.ExtID = ext_id.String
		p.BaseProductExtID = base_product_ext_id.String
		p.DurationDays = duration_days.Int32
		p.Visibility = entity.ProductVisibility(visibility.String)
		p.MaxAttendees = max_attendees.Int32
		p.Format = entity.ProductFormat(format.String)

		products = append(products, &p)
	}

	return products, nil
}
