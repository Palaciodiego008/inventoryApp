package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// Line model struct.go
type Line struct {
	ID               uuid.UUID `json:"id" db:"id"`
	PhoneLine        string    `json:"phone_line" db:"phone_line"`
	Carrier          string    `json:"carrier" db:"carrier"`
	AssociatedTo     string    `json:"associated_to" db:"associated_to"`
	AssociatedDevice string    `json:"associated_device" db:"associated_device"`
	EndContractDate  string    `json:"end_contract_date" db:"end_contract_date"`
	UpgradeEligibity time.Time `json:"upgrade_eligibity" db:"upgrade_eligibity"`
	Status           string    `json:"status" db:"status"`
	Iccid            string    `json:"iccid" db:"iccid"`
	UserID           uuid.UUID `json:"user_id" db:"user_id"`
	User             User      `json:"user,omitempty" belongs_to:"user" fk_id:"UserID"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

// Lines array model struct of Line.
type Lines []Line

func (l Line) String() string {
	jl, err := json.Marshal(l)
	if err != nil {
		fmt.Printf("error marshalling json on string Method: %v\n", err)
	}

	return string(jl)
}

// func (l Line) SelectValueU() interface{} {
// 	return l.UserID
// }

// // SelectLabel implements the selectable interface.
// func (u User) SelectLabelU() string {
// 	return u.FirstName + " " + u.LastName
// }

func (l *Line) Validate() *validate.Errors {
	return validate.Validate(
		&validators.StringIsPresent{Field: l.PhoneLine, Name: "PhoneLine", Message: "This can not be empty"},
		&validators.StringIsPresent{Field: l.Carrier, Name: "Carrier", Message: "This can not be empty"},
		&validators.StringIsPresent{Field: l.AssociatedTo, Name: "AssociatedTo", Message: "This can not be empty"},
		&validators.StringIsPresent{Field: l.AssociatedDevice, Name: "AssociatedDevice", Message: "This can not be empty"},
		&validators.StringIsPresent{Field: l.EndContractDate, Name: "EndContractDate", Message: "This can not be empty"},
		&validators.StringIsPresent{Field: l.Status, Name: "Status", Message: "This can not be empty"},
		&validators.StringIsPresent{Field: l.Iccid, Name: "Iccid", Message: "This can not be empty"},
	)
}
