package repository

import (
	"log"

	"github.com/dgraph-io/badger"
	c "github.com/guard-trader/config"
	ent "github.com/guard-trader/entities"
	"github.com/guard-trader/helpers"
)

// Badger db
type Badger struct {
	DB *badger.DB
}

// OpenBadger create instance of badger
func OpenBadger() *Badger {
	path := c.Config.Database.Path
	opts := badger.DefaultOptions
	opts.Dir = path
	opts.ValueDir = path
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	return &Badger{DB: db}
}

// GetAccounts accounts of guard
func (bg Badger) GetAccounts() ([]ent.Account, error) {
	accs := []ent.Account{}
	err := bg.DB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			v, err := item.Value()
			if err != nil {
				return err
			}
			accs = append(accs, helpers.DeserializerAccount(v))
		}
		return nil
	})
	return accs, err
}

// RegisterAccount set new account
func (bg Badger) RegisterAccount(acc ent.Account) error {
	err := bg.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set(helpers.SerializerString(acc.ID), helpers.SerializerAccount(acc))
		return err
	})
	return err
}
