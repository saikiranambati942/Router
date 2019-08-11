package app

import (
	"Router/config"
	"log"
)

//Start ... It is the main startup function for the app package
func Start() error {
	log.Println("Router application started")
	var rConfig = &config.Routes{}
	//read config
	err := config.LoadConfig(rConfig)
	if err != nil {
		return err
	}
	//Init passes config details to routing go file
	Init(rConfig)
	return err
}
