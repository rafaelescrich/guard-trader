package repository

import (
	"fmt"

	"github.com/boltdb/bolt"
	acc "github.com/guard-trader/account"
	c "github.com/guard-trader/config"
	hlp "github.com/guard-trader/helpers"
)

// PutUpdateAccount insert acc
func (db *BoltDB) PutUpdateAccount(account acc.Account) error {
	err := db.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(c.Config.Buckets.Account))
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = b.Put(hlp.SerializerString(account.ID), hlp.SerializerAccount(account))
		if err != nil {
			fmt.Println("error on insert manager db ", err)
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// GetAccount search for an acc
func (db *BoltDB) GetAccount(ID string) (acc.Account, error) {
	account := acc.Account{}
	err := db.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(c.Config.Buckets.Account))
		result := b.Get(hlp.SerializerString(ID))
		if result == nil {
			return fmt.Errorf("account not inserted")
		}
		account = hlp.DeserializerAccount(result)
		return nil
	})
	return account, err
}

// DeleteAccount delete an acc
func (db *BoltDB) DeleteAccount() error {
	return nil
}
