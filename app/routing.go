package app

import (
	"Router/config"
	"io"
	"net/http"
	"strings"
)

var conf *config.Routes

//Foo ... It is a handler to handle the http requests
func Foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("input")
	io.WriteString(w, "Input request is  : "+v)
	io.WriteString(w, "\nRouting to server : "+FindRoute(v))
}

//FindRoute ... takes a string as an input and returns the server name
func FindRoute(input string) string {
	var server string
	var sourceLength = -1
	//looping over the routes
	for _, v := range conf.Paths {
		//Removing the `.*` from the config route
		source := strings.TrimRight(v.Source, ".*")
		//checking whether the input string passed matches or sub matches with particular route
		bool := strings.Contains(input, source)
		// looping over matched routes with the input to get the maximum matched route to find the server to the input request
		if bool {
			if len(source) > sourceLength {
				server = v.Destination
				sourceLength = len(source)
			}
		}

	}
	return server
}

//Init initializes the variable with config details
func Init(config *config.Routes) {
	conf = config
}
