package app

import (
	"log"
	"net/http"

	"github.com/eliudarudo/microservices/level1/controllers"
	"github.com/gorilla/mux"
)

// App to create Router Instance
type App struct {
	Router *mux.Router
}

// Initialize initializes the Router Instance
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets all the required routers
func (a *App) setRouters() {
	// Index route
	a.Get("/", a.handleRequest(controllers.Index))
	a.Get("/get-json", a.handleRequest(controllers.GetJSON))

	a.Put("/db", a.handleRequest(controllers.PutNewItem))
	a.Get("/db", a.handleRequest(controllers.GetItemHandler))
	a.Post("/db", a.handleRequest(controllers.ModifyItemHandler))
	a.Delete("/db", a.handleRequest(controllers.DeleteItemHandler))
	a.Get("/db/all", a.handleRequest(controllers.GetItemsHandler))
	a.Delete("/db/all", a.handleRequest(controllers.DeleteAllItemsHandler))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

// RequestHandlerFunction does something
type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
