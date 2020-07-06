package app

import (
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestFindRoute(t *testing.T) {
	os.Setenv("ConfigFile", os.Getenv("GOPATH")+"\\src\\"+"Router\\config.json")
	Start()
	time.Sleep(1 * time.Second)
	route := FindRoute("customer1.us.ca")
	if route != "server1" {
		log.Fatal("error occurred")
	}
}
func TestFoo(t *testing.T) {
	os.Setenv("ConfigFile", os.Getenv("GOPATH")+"\\src\\"+"Router\\config.json")
	Start()
	req := httptest.NewRequest("GET", "http://localhost:8080/?input=customer1.us.ca.sfo", nil)
	w := httptest.NewRecorder()
	Foo(w, req)
	fmt.Printf("%d - %s\n", w.Code, w.Body.String())
	if w.Code != 200 {
		log.Fatal("error occurred")
	}
}
func TestInit(t *testing.T) {
	os.Setenv("ConfigFile", os.Getenv("GOPATH")+"\\src\\"+"Router\\config.json")
	Start()
	if conf == nil {
		log.Fatal("error occurred")
	}
}
