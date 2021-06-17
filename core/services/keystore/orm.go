package keystore

import (
	"context"
	"math/rand"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/csakey"
	"github.com/smartcontractkit/chainlink/core/services/keystore/keys/ethkey"
	"gorm.io/gorm"
)

func NewORM(db *gorm.DB) ksORM {
	return ksORM{
		db: db,
	}
}

type ksORM struct {
	db *gorm.DB
}

func (orm ksORM) saveEncryptedKeyRing(kr *encryptedKeyRing) error {
	err := orm.db.Exec(`
		UPDATE encrypted_key_rings
		SET encrypted_keys = ?
		WHERE id = (
			SELECT id FROM encrypted_key_rings ORDER BY id ASC LIMIT 1
		)
	`, kr.EncryptedKeys).Error
	if err != nil {
		return errors.Wrap(err, "while saving keyring")
	}
	return nil
}

func (orm ksORM) getEncryptedKeyRing() (encryptedKeyRing, error) {
	kr := encryptedKeyRing{}
	err := orm.db.First(&kr).Error
	if err == gorm.ErrRecordNotFound {
		kr = encryptedKeyRing{ID: rand.Int31()}
		err2 := orm.db.Create(&kr).Error
		if err2 != nil {
			return kr, err2
		}
	} else if err != nil {
		return kr, err
	}
	return kr, nil
}
func (orm ksORM) getEthKeyStateWhere(query string, args ...interface{}) (state ethkey.State, _ error) {
	return state, orm.db.Where(query, args...).First(&state).Error
}

func (orm ksORM) getEthKeyStatesWhere(query string, args ...interface{}) (states []ethkey.State, _ error) {
	return states, orm.db.Where(query, args...).Find(&states).Error
}

func (orm ksORM) getNextRoundRobinAddress(whitelist []common.Address) (ethkey.State, error) {
	var query *gorm.DB
	if len(whitelist) == 0 {
		query = orm.db
	} else {
		query = orm.db.Where("address in ?", whitelist)
	}
	var state ethkey.State
	err := query.
		Where("is_funding = false").
		Order("last_used ASC").
		First(&state).
		Error
	if err != nil {
		return ethkey.State{}, err
	}
	return state, nil
}

// ~~~~~~~~~~~~~~~~~~~~ LEGACY FUNCTIONS FOR V1 MIGRATION ~~~~~~~~~~~~~~~~~~~~

// GetEncryptedV1CSAKeys lists all csa keys.
func (o ksORM) GetEncryptedV1CSAKeys(ctx context.Context) ([]csakey.Key, error) {
	keys := []csakey.Key{}
	stmt := `
		SELECT id, public_key, encrypted_private_key, created_at, updated_at
		FROM csa_keys;
	`

	err := o.db.Raw(stmt).Scan(&keys).Error
	if err != nil {
		return keys, err
	}

	return keys, nil
}
