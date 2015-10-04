# simple-server-watch-openfd-s
this application is just a simple server that will not quit when the open file idscriptor reached ..
it simply waits for the other connection to get closed and then allow others to connect to it 

the client test the maximum number of connections in 1 second ... 

you should simply run this app like below 

`go run server.go`

and after that 

`go run client.go`
