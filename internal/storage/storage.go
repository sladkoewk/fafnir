package storage

import (
	"errors"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/fafnir/internal/storage/boltdb"
)

type AppStorage struct {
	db *bolt.DB
}

func InitBolt() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Batch(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(boltdb.Sheet))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(boltdb.Email))
		return err
	}); err != nil {
		return nil, err
	}

	return db, nil
}

func NewStorage(db *bolt.DB) *AppStorage {
	return &AppStorage{db: db}
}

func (s *AppStorage) Save(id int64, value string, bucket boltdb.Bucket) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intToBytes(id), []byte(value))
	})
}

func (s *AppStorage) Get(id int64, bucket boltdb.Bucket) (string, error) {
	var token string

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		token = string(b.Get(intToBytes(id)))
		return nil
	})

	if token == "" {
		return "", errors.New("not found")
	}

	return token, err
}

func intToBytes(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
