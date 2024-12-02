package entity

type Token struct {
	AuthToken string `json:"access_token"`
}

func (t *Token) Validate() error {
	if t.AuthToken == "" {
		return ErrInvalidEntity
	}
	return nil
}
