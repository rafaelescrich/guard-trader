package db

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

// BoltDB model
type BoltDB struct {
	DB *bolt.DB
}

// TODO: Orders history, ranks, balance

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
	return db
}
