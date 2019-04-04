package migrations

import (
	"github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

type Account201904030129 struct {
	ID                string `gorm:"index;primary_key"`
	UserID            string `gorm:"index"`
	Email             string `gorm:"index"`
	Password          string
	PasswordResetCode string
	Type              string `gorm:"index"`
	Verified          bool   `gorm:"index"`
	CreatedAt         uint32 `gorm:"index"`
	UpdatedAt         uint32 `gorm:"index"`
}

type User201904030129 struct {
	ID          string `gorm:"index;primary_key"`
	FirstName   string `gorm:"index"`
	LastName    string `gorm:"index"`
	Deactivated bool   `gorm:"index"`
	CreatedAt   uint32 `gorm:"index"`
	UpdatedAt   uint32 `gorm:"index"`
}

func init() {
	migration := gormigrate.Migration{
		ID: "201904030129",

		Migrate: func(tx *gorm.DB) error {

			type Account struct {
				Account201904030129
			}

			type User struct {
				User201904030129
			}

			return tx.
				CreateTable(&Account{}).
				CreateTable(&User{}).
				Error
		},

		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	}

	RegisterMigration(&migration)
}
