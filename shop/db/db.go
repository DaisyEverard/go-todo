package db

import (
	"errors"
)

type Item struct {
	name string
	price float64
}

type DataBase struct {
	items map[int]*Item
}

func GetNewDB() *DataBase {
	return &DataBase{items: make(map[int]*Item)}
}

func (db *DataBase) FindItem(id int) (*Item, error) {
	item, exists := db.items[id]
	if exists {
		return item, nil
	} else {
		return nil, errors.New("no item found")
	}
}

func (db *DataBase) AddItem(id int, name string, price float64) {
	newItem := &Item{name: name, price: price}
	db.items[id] = newItem
}
