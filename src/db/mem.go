package db

import "caiomcg.com/playing_cards/src/models"
import "errors"
import "sync"

type Database struct {
	db []models.Deck
}

var (
	instance    *Database
	intanceOnce sync.Once
)

func (d *Database) Find(id string) (*models.Deck, error) {
	for i, _ := range d.db {
		if d.db[i].Id.String() == id {
			return &d.db[i], nil
		}
	}

	return &models.Deck{}, errors.New("Deck not found")
}

func (d *Database) Insert(deck models.Deck) {
	d.db = append(d.db, deck)
}

func (d *Database) GetAll() []models.Deck {
	return d.db
}

func (d *Database) Wipe() {
	d.db = []models.Deck{}
}

func (d *Database) Peek() (models.Deck, error) {
	if len(d.db) == 0 {
		return models.Deck{}, errors.New("No content in DB")
	}

	return d.db[0], nil
}

func Instance() *Database {
	intanceOnce.Do(func() {
		instance = &Database{db: []models.Deck{}}
	})
	return instance
}
