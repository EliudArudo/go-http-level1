package store

import (
	"errors"
)

// Item sends objects back
type Item struct {
	ID   string `json:"id"`
	Item string `json:"item"`
}

// ItemResponseStatus sends simple response about object without item.value
type ItemResponseStatus struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type SimplestResponseStatus struct {
	Status string `json:"status"`
}

var dataStore = make(map[string]string)

// AddItem gets body from 'PUT' request and places in our map
func AddItem(id string, item string) ItemResponseStatus {
	dataStore[id] = item

	return ItemResponseStatus{ID: id, Status: "OK"}
}

// GetItem takes in id from controller and returns the item found
func GetItem(id string) (Item, error) {
	value, found := dataStore[id]
	if !found { // If not found
		return Item{}, errors.New("Item not found")
	}

	return Item{ID: id, Item: value}, nil
}

// GetItems gets all we have in the dataStore
func GetItems() []Item {

	items := []Item{}

	for k, v := range dataStore {
		items = append(items, Item{ID: k, Item: v})
	}

	return items
}

// ModifyItem changes item value in dataStore
func ModifyItem(id string, item string) (ItemResponseStatus, error) {
	_, found := dataStore[id]
	if !found {
		return ItemResponseStatus{}, errors.New("Object non existent in our database")
	}
	dataStore[id] = item

	return ItemResponseStatus{ID: id, Status: "Successfully modified"}, nil
}

// DeleteItem deletes item from dataStore
func DeleteItem(id string) (ItemResponseStatus, error) {
	_, found := dataStore[id]
	if !found {
		return ItemResponseStatus{}, errors.New("Item does not exist in our store")
	}

	delete(dataStore, id)
	return ItemResponseStatus{ID: id, Status: "Successfully deleted item"}, nil
}

// DeleteAllItems clears all data from our dataStore map
func DeleteAllItems() {

	for k := range dataStore {
		delete(dataStore, k)
	}
}
