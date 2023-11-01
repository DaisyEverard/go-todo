package main

import (
	"fmt"
	shop "store/shop"
)

func main() {
	myStore := shop.NewStore("myStore")
	myStore.AddItem(5, "home alone", 20.5)

	foundItem, err := myStore.FindItem(5)
	fmt.Println(foundItem, err)

	foundItem, err = myStore.FindItem(2)
	fmt.Println(foundItem, err)
}
