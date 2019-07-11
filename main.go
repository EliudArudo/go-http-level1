package main

import (
	"fmt"

	"github.com/eliudarudo/microservices/level1/app"
)

func main() {
	port := 8080
	// portString := ":" + string(port)
	portString := fmt.Sprintf(":%d", port)

	app := &app.App{}
	app.Initialize()
	fmt.Printf("Starting server on port: %v\n", port)
	app.Run(portString)

}
