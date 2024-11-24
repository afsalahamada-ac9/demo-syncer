/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

import (
	"time"
)

// Course type
type CourseType string

const (
	CourseInPerson CourseType = "in-person"
	CourseOnline   CourseType = "online"
	// Add new types here
)

// Course status
type CourseStatus string

const (
	CourseDraft            CourseStatus = "draft"
	CourseArchived         CourseStatus = "archived"
	CourseOpen             CourseStatus = "open"
	CourseExpenseSubmitted CourseStatus = "expense-submitted"
	CourseExpenseDeclined  CourseStatus = "expense-declined"
	CourseClosed           CourseStatus = "closed"
	CourseActive           CourseStatus = "active"
	CourseDeclined         CourseStatus = "declined"
	CourseSubmitted        CourseStatus = "submitted"
	CourseCanceled         CourseStatus = "canceled"
	CoursedInactive        CourseStatus = "inactive"
	// Add new types here
)

// Course Location
// TODO: json tags must be moved to presenter
type CourseLocation struct {
	Street1 string `json:"street"`
	Street2 string `json:"street_2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

// Course date/time
type CourseDateTime struct {
	Date      string `json:"date"`      // Only date in YYYY-MM-DD format
	StartTime string `json:"startTime"` // Only time in HH:MM:SS format (SS is optional, default 00)
	EndTime   string `json:"endTime"`
}

// Course data
type Course struct {
	ID       ID
	TenantID ID
	CenterID ID

	ExtID string

	Name     string
	Notes    string
	Timezone string

	Location CourseLocation

	Status CourseStatus

	// TODO: CType to be renamed to delivery mode or method
	CType CourseType

	MaxAttendees int32
	NumAttendees int32

	// TODO: AutoApprove may not be required here. It's likely in Course Master or Catalog.
	IsAutoApprove bool

	// meta data
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewCourseLocation creates a new course location
func NewCourseLocation(street1 string,
	street2 string,
	city string,
	state string,
	zip string,
	country string) (*CourseLocation, error) {

	l := &CourseLocation{
		Street1: street1,
		Street2: street2,
		City:    city,
		State:   state,
		Zip:     zip,
		Country: country,
	}
	err := l.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return l, nil
}

// Validate validates course location
func (l *CourseLocation) Validate() error {
	if l.Street1 == "" || l.City == "" || l.State == "" || l.Zip == "" || l.Country == "" {
		return ErrInvalidEntity
	}
	return nil
}

// NewCourse create a new course
func NewCourse(tenantID ID,
	extID string,
	centerID ID,
	name string,
	notes string,
	timezone string,
	location CourseLocation,
	status CourseStatus,
	ctype CourseType,
	maxAttendees int32,
	numAttendees int32,
	isAutoApprove bool) (*Course, error) {
	c := &Course{
		ID:            NewID(),
		TenantID:      tenantID,
		ExtID:         extID,
		CenterID:      centerID,
		Name:          name,
		Notes:         notes,
		Timezone:      timezone,
		Location:      location,
		Status:        status,
		CType:         ctype,
		MaxAttendees:  maxAttendees,
		NumAttendees:  numAttendees,
		IsAutoApprove: isAutoApprove,
		CreatedAt:     time.Now(),
	}
	err := c.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return c, nil
}

// Validate validate course
func (c *Course) Validate() error {
	if c.Name == "" {
		return ErrInvalidEntity
	}
	return nil
}
