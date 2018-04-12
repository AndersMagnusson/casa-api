package database

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func SetDatabase(database *bolt.DB) {
	db = database
}

func CreateBucketIfNotExists(name string) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return fmt.Errorf("Failed to get bucket: %s, err: %s", name, err.Error())
		}
		return nil
	})
}

func View(f func(tx *bolt.Tx) error) error {
	return db.View(f)
}

func Update(f func(tx *bolt.Tx) error) error {
	return db.Update(f)
}

func Put(bucketName string, data interface{}) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		encoded, err := json.Marshal(data)
		if err != nil {
			return err
		}
		b.Put([]byte(bucketName), encoded)
		return nil
	})
}
