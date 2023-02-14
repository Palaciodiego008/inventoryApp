package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Line2 struct {
	ID               uuid.UUID `json:"id" db:"id"`
	PhoneLine        string    `json:"phone_line" db:"phone_line"`
	Carrier          string    `json:"carrier" db:"carrier"`
	AssociatedTo     string    `json:"associated_to" db:"associated_to"`
	AssociatedDevice string    `json:"associated_device" db:"associated_device"`
	EndContractDate  string    `json:"end_contract_date" db:"end_contract_date"`
	UpgradeEligibity time.Time `json:"upgrade_eligibity" db:"upgrade_eligibity"`
	Status           string    `json:"status" db:"status"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

func (ms *ModelSuite) Test_Line() {
	id, err := uuid.FromString("468d02bb-98ac-4496-af51-62c3c1f55530")
	ms.NoError(err)
	newLine := Line{
		ID:               id,
		Carrier:          "Carrier Test",
		AssociatedTo:     "Test",
		AssociatedDevice: "Mobile test",
		EndContractDate:  "12-22-09",
		UpgradeEligibity: time.Now(),
		Status:           "status test",
	}

	ms.NoError(ms.DB.Create(&newLine))
}
