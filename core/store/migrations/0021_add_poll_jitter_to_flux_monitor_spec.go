package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

const (
	up21 = `
ALTER TABLE flux_monitor_specs
ADD poll_jitter integer DEFAULT 0;
`
	down21 = `
ALTER TABLE flux_monitor_specs
DROP COLUMN poll_jitter;
`
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "0021_add_poll_jitter_to_flux_monitor_spec",
		Migrate: func(db *gorm.DB) error {
			return db.Exec(up21).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Exec(down21).Error
		},
	})
}
