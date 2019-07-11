package logging

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eliudarudo/microservices/level1/store"
)

// LogRequest gives us query params and body of whatever's coming in
func LogRequest(r *http.Request, body ...interface{}) {
	route := r.RequestURI
	method := r.Method
	queryParams := r.URL.Query().Get("id") // specific at this point

	item := store.Item{} // What we're expecting to come in

	fmt.Println("-------------- What we got----------------") // Because we'll always be starting
	fmt.Printf("Route requested: %s", route)
	fmt.Println()
	fmt.Printf("Method is: %s", method)

	fmt.Println()
	if len(queryParams) > 0 {
		fmt.Printf("Query params, id: %s", queryParams)
	}
	fmt.Println()
	if len(body) > 0 {
		marshalledItem, _ := json.Marshal(item)
		fmt.Println("Body params: ")
		fmt.Printf("%s", marshalledItem)
		fmt.Println()
	}

}

// LogResponse logs what's being sent back to the user
func LogResponse(response string) {
	fmt.Println("-------------- What we're sending out -----------------")
	fmt.Printf("%v", string(response))
	fmt.Println()
	fmt.Println("--------------------------------------")
}
