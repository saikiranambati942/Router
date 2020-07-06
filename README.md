# Router

This application loads the config file during initialization and finds the route to a appropriate destination server based on the config details loaded at the startup i.e; it will route to destination server based on best/maximum matched input request.

Q1) How to run the router application?

A1) Run the command `go run main.go` in the terminal.

Q2) How to send request to the router application?

A2) After started running the application and making sure that the application is up, go to system browser and make a request in this pattern below
`http://localhost:8080/?input=customer1.us.ca.sfo`
Here ":8080" is the exposed port for the http server we used in the application and send the request in the query parameter(input).

Sample o/p for the above request:

`Input request is  : customer1.us.ca.sfo`

`Routing to server : server1 `

Q3) Want to see documentation of the router application?

A3) Run the command ` godoc -http=":6060"` in the terminal and make a request in the browser this way `localhost:6060`.

`Note: Go documentation even contains all packages including standard libraries`


# DockerImage : `docker pull saikiran942/routerimage`

docker run -d $containername -p $host port:$application port

example : docker run -p 8081:8080 -d $imagename




 
 
 
