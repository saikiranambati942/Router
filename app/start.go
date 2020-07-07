package app

import (
	"Router/config"
	"fmt"
	"log"
	"time"
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
	Init(rConfig)
	configUpdate(rConfig)
	return err
}

func configUpdate(rConfig *config.Routes) {
	// we can set the ticker value as environment variable
	ticker := time.NewTicker(20 * time.Second)
	go func() {
		for {
			<-ticker.C
			fmt.Println("Loading configuration Update...")
			config.LoadConfig(rConfig)

		}
	}()
}
