package migrations

import (
	"gorm.io/gorm"
)

const up45 = `
ALTER TABLE eth_txes ADD COLUMN access_list jsonb;
ALTER TABLE eth_tx_attempts
	ADD COLUMN tx_type smallint NOT NULL DEFAULT 0,
	ADD COLUMN gas_tip_cap numeric(78,0),
	ADD COLUMN gas_fee_cap numeric(78,0),
	ADD CONSTRAINT chk_tx_type_is_byte CHECK (
		tx_type >= 0 AND tx_type <= 255
	),
	ADD CONSTRAINT chk_legacy_or_dynamic CHECK (
		(tx_type = 0 AND gas_price IS NOT NULL AND gas_tip_cap IS NULL AND gas_fee_cap IS NULL) 
		OR
		(tx_type = 2 AND gas_price IS NULL AND gas_tip_cap IS NOT NULL AND gas_fee_cap IS NOT NULL)
	)
;
`

const down45 = `
ALTER TABLE eth_txes DROP COLUMN access_list;
ALTER TABLE eth_tx_attempts DROP COLUMN tx_type, DROP COLUMN gas_tip_cap, DROP COLUMN gas_fee_cap;
`

func init() {
	Migrations = append(Migrations, &Migration{
		ID: "0045_dynamic_fee_txes",
		Migrate: func(db *gorm.DB) error {
			return db.Exec(up45).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Exec(down45).Error
		},
	})
}
