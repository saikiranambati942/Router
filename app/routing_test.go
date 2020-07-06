package app

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestFindRoute(t *testing.T) {
	os.Setenv("ConfigFile",os.Getenv("GOPATH")+"\\src\\"+"Router\\config.json")
	Start()
	time.Sleep(1*time.Second)
	route:=FindRoute("customer1.us.ca")
	if route!="server1"{
		log.Fatal("error occurred")
	}
}
