package utils

import (
	"github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/couchdb/v4"
	"context"
)

var db *kivik.DB

func ConnectDB() (*kivik.DB, error) {
	client, err := kivik.New("couch", "http://admin:adminpw@localhost:5984/")
	if err != nil {
		return nil, err
	}

	// Check if the database exists
	exists, err := client.DBExists(context.Background(),"poc")
	if err != nil {
		return nil, err
	}

	if !exists {
		// Create the "poc" database
		err = client.CreateDB(context.Background(),"poc")
		if err != nil {
			return nil, err
		}
	}

	db = client.DB("poc")

	return db, nil
}

func GetDB() *kivik.DB {
	return db
}
