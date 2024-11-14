package entity

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Tenant data
type Tenant struct {
	ID       ID
	Username string // same as email
	Password string

	AuthToken string

	// meta data
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewTenant create a new tenant
func NewTenant(username, password string) (*Tenant, error) {
	t := &Tenant{
		ID:        NewID(),
		Username:  username,
		Password:  genPassword(password),
		CreatedAt: time.Now(),
	}
	err := t.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return t, nil
}

// Validate validate tenant
func (t *Tenant) Validate() error {
	if t.Password == "" {
		log.Printf("Invalid password")
		return ErrInvalidEntity
	}

	if t.Username == "" {
		return ErrInvalidEntity
	}
	return nil
}

// ValidatePassword validate tenant password
func (t *Tenant) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(t.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}

func genPassword(raw string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		log.Printf("Unable to generate password, %v", err.Error())
		return ""
	}

	return string(hash)
}

// ValidateToken validate tenant auth token
func (t *Tenant) ValidateToken(token string) error {
	if token != t.genToken() {
		return ErrTokenMismatch
	}

	return nil
}

// simple token generator for now
func (t *Tenant) genToken() string {
	token := sha256.New()
	token.Write([]byte(t.Password))
	hash := token.Sum(nil)

	return base64.StdEncoding.EncodeToString(hash)
}

// Exported API to generate and store the token
func (t *Tenant) GenToken() error {
	t.AuthToken = t.genToken()
	return nil
}
