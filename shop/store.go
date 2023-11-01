package shop

import (
    "store/shop/db"
)

type Store struct {
	StoreName string
	priceDB *db.DataBase
}

func NewStore(name string) *Store {
	return &Store{StoreName: name, priceDB: db.GetNewDB()}
}

func (store *Store) FindItem(id int) (*db.Item, error) {
	return store.priceDB.FindItem(id)
}
func (store *Store) AddItem(id int, name string, price float64){
	store.priceDB.AddItem(id, name, price)
}