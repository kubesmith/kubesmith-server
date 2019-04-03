package migrations

import (
	"github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

type Migrations []*gormigrate.Migration

var migrations Migrations

func (m Migrations) Len() int {
	return len(m)
}

func (m Migrations) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m Migrations) Less(i, j int) bool {
	return (m[i].ID <= m[j].ID)
}

func Run(db *gorm.DB) error {
	if migrations == nil {
		return nil
	}

	return gormigrate.New(
		db,
		&gormigrate.Options{
			UseTransaction: true,
		},
		migrations,
	).Migrate()
}

func RegisterMigration(migration *gormigrate.Migration) {
	if migrations == nil {
		migrations = Migrations{}
	}

	migrations = append(migrations, migration)
}
