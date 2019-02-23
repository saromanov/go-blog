// Package user defines inner representation of the user
// models, db handling
package user

// User creates db model for users
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`

	PasswordHash []byte `json:"-"`
}
