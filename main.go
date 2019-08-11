package main

import (
	"Router/app"
	"log"
	"net/http"
	"os"
)

func main() {
	//Calling the main startup function
	err := app.Start()
	if err != nil {
		log.Fatal("Error starting Router application ...", err)
		os.Exit(1)
	}
	//Default handler
	http.HandleFunc("/", app.Foo)
	//ListenAndServe starts an HTTP server with a given address and handler.The handler is usually nil, which means to use default handler HandleFunc
	http.ListenAndServe(":8080", nil)
}
