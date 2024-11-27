/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package repository

import (
	"database/sql"
	"encoding/json"
	"time"

	"sudhagar/glad/entity"
)

// CoursePGSQL mysql repo
type CoursePGSQL struct {
	db *sql.DB
}

// NewCoursePGSQL create new repository
func NewCoursePGSQL(db *sql.DB) *CoursePGSQL {
	return &CoursePGSQL{
		db: db,
	}
}

// Create creates a course
func (r *CoursePGSQL) Create(e *entity.Course) (entity.ID, error) {
	locationJSON, err := json.Marshal(e.Location)
	if err != nil {
		return e.ID, err
	}

	stmt, err := r.db.Prepare(`
		INSERT INTO course (id, tenant_id, ext_id, center_id, name, notes, timezone, location, status, ctype, max_attendees, num_attendees, is_auto_approve, created_at) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.TenantID,
		e.ExtID,
		e.CenterID,
		e.Name,
		e.Notes,
		e.Timezone,
		string(locationJSON),
		e.Status,
		e.CType,
		e.MaxAttendees,
		e.NumAttendees,
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

// Get retrieves a course
func (r *CoursePGSQL) Get(id entity.ID) (*entity.Course, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, tenant_id, ext_id, center_id, name, notes, timezone, location,
		status, ctype, max_attendees, num_attendees, is_auto_approve, created_at
		FROM course
		WHERE id = $1;`)
	if err != nil {
		return nil, err
	}
	var c entity.Course
	var ext_id sql.NullString
	var name, notes, timezone, loc_json, status, ctype sql.NullString
	err = stmt.QueryRow(id).Scan(&c.ID, &c.TenantID, &ext_id, &c.CenterID, &name, &notes, &timezone,
		&loc_json, &status, &ctype, &c.MaxAttendees, &c.NumAttendees, &c.IsAutoApprove, &c.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if loc_json.Valid && loc_json.String != "" {
		err = json.Unmarshal([]byte(loc_json.String), &c.Location)
		if err != nil {
			return nil, err
		}
	}

	c.ExtID = ext_id.String
	c.Name = name.String
	c.Notes = notes.String
	c.Timezone = timezone.String
	c.Status = entity.CourseStatus(status.String)
	c.CType = entity.CourseType(ctype.String)

	return &c, nil
}

// Update updates a course
func (r *CoursePGSQL) Update(e *entity.Course) error {
	e.UpdatedAt = time.Now()
	locationJSON, err := json.Marshal(e.Location)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`
		UPDATE course SET center_id = $1, name = $2, notes = $3, timezone = $4, location = $5,
		status = $6, ctype = $7, max_attendees = $8, num_attendees = $9, is_auto_approve = $10,
		updated_at = $11
		WHERE id = $12;`,
		e.CenterID, e.Name, e.Notes, e.Timezone, string(locationJSON), (e.Status), (e.CType),
		e.MaxAttendees, e.NumAttendees, e.IsAutoApprove, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

// Search searches courses
func (r *CoursePGSQL) Search(tenantID entity.ID,
	q string, page, limit int,
) ([]*entity.Course, error) {
	query := `
		SELECT id, tenant_id, ext_id, center_id, name, notes, timezone, location,
		status, ctype, max_attendees, num_attendees, is_auto_approve, created_at
		FROM course
		WHERE tenant_id = $1 AND name LIKE $2`

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

// List lists courses
func (r *CoursePGSQL) List(tenantID entity.ID, page, limit int) ([]*entity.Course, error) {
	query := `
		SELECT id, tenant_id, ext_id, center_id, name, notes, timezone, location,
		status, ctype, max_attendees, num_attendees, is_auto_approve, created_at
		FROM course
		WHERE tenant_id = $1`

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

// Delete deletes a course
func (r *CoursePGSQL) Delete(id entity.ID) error {
	res, err := r.db.Exec(`DELETE FROM course WHERE id = $1;`, id)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Get total courses
func (r *CoursePGSQL) GetCount(tenantID entity.ID) (int, error) {
	stmt, err := r.db.Prepare(`SELECT count(*) FROM course WHERE tenant_id = $1;`)
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

func (r *CoursePGSQL) scanRows(rows *sql.Rows) ([]*entity.Course, error) {
	var courses []*entity.Course
	for rows.Next() {
		var course entity.Course
		var ext_id, name, notes, timezone, is_auto_approve, created_at sql.NullString
		var max_attendees, num_attendees sql.NullInt32
		err := rows.Scan(
			&course.ID,
			&course.TenantID,
			&course.CenterID,
			&ext_id,
			&name,
			&notes,
			&timezone,
			&course.Location,
			&course.Status,
			&course.CType,
			&max_attendees,
			&num_attendees,
			&is_auto_approve,
			&created_at,
		)
		if err != nil {
			return nil, err
		}

		course.ExtID = ext_id.String
		course.Name = name.String
		course.Notes = notes.String
		course.Timezone = timezone.String
		course.NumAttendees = num_attendees.Int32
		course.MaxAttendees = max_attendees.Int32

		courses = append(courses, &course)
	}
	return courses, nil
}
