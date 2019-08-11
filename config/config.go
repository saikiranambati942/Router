package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

//Route struct to handle route config
type Route struct {
	Source      string `json:"route"`
	Destination string `json:"server"`
}

//Routes struct defines the routes for the http calls
type Routes struct {
	Paths []Route `json:"routes"`
}

//LoadConfig ... loads the config details during the application startup
func LoadConfig(Config *Routes) (err error) {
	log.Println("Initializing config. reading from config.json")
	file, err := os.Open("config.json")
	if err != nil {
		log.Println("Couldn't able to open the config file")
		return err
	}
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Couldn't able to read the config file")
		return err
	}
	err = json.Unmarshal(fileData, Config)
	if err != nil {
		log.Println("Couldn't able to parse the config file")
		return err
	}
	log.Println("Loading config values")
	for _, v := range Config.Paths {
		log.Printf("route : %s --> servername: %s\n", v.Source, v.Destination)
	}
	return err
}
