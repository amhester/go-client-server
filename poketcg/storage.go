package main

import (
	"encoding/json"

	"github.com/boltdb/bolt"
)

var (
	deckBucket = []byte("decks")
)

type LocalStorage struct {
	dbFile string
	db     *bolt.DB
}

func NewLocalStorage(file string) (*LocalStorage, error) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(deckBucket)
		if err != nil {
			return err
		}
		return nil
	})
	return &LocalStorage{
		dbFile: file,
		db:     db,
	}, nil
}

func (ls *LocalStorage) Close() {
	ls.db.Close()
}

func (ls *LocalStorage) AddDeck(d *Deck) error {
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}
	err = ls.db.Update(func(tx *bolt.Tx) error {
		decks := tx.Bucket(deckBucket)
		err := decks.Put([]byte(d.Name), b)
		return err
	})
	return err
}
