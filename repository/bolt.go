package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
	c "github.com/guard-trader/config"
)

// BoltDB model
type BoltDB struct {
	DB *bolt.DB
}

// NewBoltDB return stance of boltdb
func NewBoltDB(subfolder string) *BoltDB {
	path := c.Config.Database.Path
	folder := path + "/guard.db"
	if subfolder != "" {
		folder = path + "/" + subfolder + "/guard.db"
	}
	if _, err := os.Stat(path + "/" + subfolder); os.IsNotExist(err) {
		_ = os.Mkdir(path+"/"+subfolder, 0771) //create if not exist
	}

	db, err := bolt.Open(folder, 0600, nil)
	if err != nil {
		fmt.Printf("BoltDB Error: %s \r\n", err)
		os.Exit(1)
	}

	return &BoltDB{db}
}

// Start start module test database
func (db *BoltDB) Start() *BoltDB {
	db.createManagerBucket()
	db.createAccountBucket()
	return db
}

func (db *BoltDB) createManagerBucket() error {
	err := db.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(c.Config.Buckets.Manager))
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (db *BoltDB) createAccountBucket() error {
	err := db.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(c.Config.Buckets.Account))
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
