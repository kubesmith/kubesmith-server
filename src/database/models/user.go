package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Deactivated bool   `json:"deactivated"`
	CreatedAt   uint32 `json:"createdAt"`
	UpdatedAt   uint32 `json:"updatedAt"`
}

func (u *User) BeforeCreate() error {
	u.ID = uuid.NewV4().String()
	u.CreatedAt = uint32(time.Now().Unix())

	return nil
}

func (u *User) BeforeSave() error {
	u.UpdatedAt = uint32(time.Now().Unix())

	return nil
}
