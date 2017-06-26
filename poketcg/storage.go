package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"io/ioutil"
	"regexp"

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
	db, err := bolt.Open("my.db", 0644, nil)
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

func (ls *LocalStorage) GetDeck(name string) (*Deck, error) {
	var raw []byte
	err := ls.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(deckBucket)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", deckBucket)
		}
		raw = bucket.Get([]byte(name))
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(raw) == 0 {
		return nil, errors.New("No deck found")
	}
	deck := &Deck{}
	err = json.Unmarshal(raw, deck)
	if err != nil {
		return nil, err
	}
	return deck, nil
}

func (ls *LocalStorage) ImportDeck(filePath string) error {
	extTest := regexp.MustCompile(`\.json$`)
	if extTest.Match([]byte(filePath)) == false {
		return errors.New("Can only import files of type json.")
	}
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	deck := &Deck{}
	err = json.Unmarshal(raw, deck)
	if err != nil {
		return err
	}
	err = ls.AddDeck(deck)
	if err != nil {
		return err
	}
	return nil
}
