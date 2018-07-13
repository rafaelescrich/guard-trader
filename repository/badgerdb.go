package repository

import (
	"log"

	"github.com/dgraph-io/badger"
	c "github.com/guard-trader/config"
)

// BadgerDB model
type BadgerDB struct {
	DB *badger.DB
}

func createBadger() *BadgerDB {
	path := c.Config.Database.Path
	folder := path + "/badger"
	opts := badger.DefaultOptions
	opts.Dir = folder
	opts.ValueDir = folder
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	return &BadgerDB{DB: db}
}
