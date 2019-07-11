package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/eliudarudo/microservices/level1/logging"
	"github.com/eliudarudo/microservices/level1/store"
)

type indexStruct struct {
	Path  string `json:"path"`
	Route string `json:"route"`
}

// Index controller receives GET requests from '/' route
func Index(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)
	res := indexStruct{Path: "index", Route: r.RequestURI}

	respondJSON(w, http.StatusOK, res)
}

// GetJSON simply returns JSON to see if everythings is OK
func GetJSON(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)

	response := store.SimplestResponseStatus{Status: "200 OK"}
	respondJSON(w, http.StatusOK, response)
}

// GetItemHandler gets an Item from our store using route parameters
func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)
	id := r.URL.Query().Get("id")

	item, err := store.GetItem(id)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	// If everything is OK
	respondJSON(w, http.StatusOK, item)

}

// GetItemsHandler gets all the items from our store
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)
	items := store.GetItems()

	respondJSON(w, http.StatusOK, items)
}

// PutNewItem places new item in our store
func PutNewItem(w http.ResponseWriter, r *http.Request) {
	item := store.Item{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	body, _ := json.Marshal(item)
	logging.LogRequest(r, string(body))

	ret := store.AddItem(item.ID, item.Item)

	respondJSON(w, http.StatusAccepted, ret)
	// store.AddItem()
}

// ModifyItemHandler takes data from request body and sends it through store function
func ModifyItemHandler(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)
	item := store.Item{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	body, _ := json.Marshal(item)
	logging.LogRequest(r, string(body))

	retrievedItem, err := store.ModifyItem(item.ID, item.Item)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, retrievedItem)
}

// DeleteItemHandler get's the route params and forwards to the storeHandler
func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)
	id := r.URL.Query().Get("id")

	response, err := store.DeleteItem(id)
	if err != nil {
		respondError(w, http.StatusBadGateway, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, response)
}

// DeleteAllItemsHandler fowards request to store function
func DeleteAllItemsHandler(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)
	store.DeleteAllItems()

	status := store.SimplestResponseStatus{Status: "Everything deleted successfully"}
	respondJSON(w, http.StatusOK, status)
}

// Utility functions
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	// Our own here
	logging.LogResponse(string(response))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, code int, message string) {
	logging.LogResponse(message)
	respondJSON(w, code, map[string]string{"error": message})
}
