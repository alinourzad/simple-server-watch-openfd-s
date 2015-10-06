// this is the client for the server
// this app simply test the the number of connection
// to the server

package main

import (
	"log"
	"net"
	// "encoding/gob"
)

func testConn(i int, x chan bool){
	log.Println("running : ", "---" , i, "---")
	c, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		log.Fatal("net.dial | ", err)
	}
	defer c.Close()
	log.Println("connected to : ", c.RemoteAddr())
	log.Println("we are : " , c.LocalAddr())
	x<-true
}

func client() {
	x := make(chan bool)
	log.Println("Running the App ^^")
	for i := 0; ; i++ {
		go testConn(i, x)
		<-x
	}
}
