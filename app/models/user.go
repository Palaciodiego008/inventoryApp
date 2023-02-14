package models

import (
	"time"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// User model struct.
type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Lines     []Line    `json:"lines,omitempty" has_many:"lines"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Users array model struct of User.
type Users []User

func (u *User) Validate() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.FirstName, Name: "FirstName", Message: "This can not be empty"},
		&validators.StringIsPresent{Field: u.LastName, Name: "LastName", Message: "This can not be empty"},
		&validators.StringIsPresent{Field: u.Email, Name: "Email", Message: "This can not be empty"},
	)
}

// SelectValue implements the selectable interface.
func (u User) SelectValue() interface{} {
	return u.ID
}

// SelectLabel implements the selectable interface.
func (u User) SelectLabel() string {
	return u.FirstName + " " + u.LastName
}
