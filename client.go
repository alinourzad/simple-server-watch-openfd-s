// this is the client for the server
// this app simply test the the number of connection
// to the server

package main

import (
	"log"
	"net"
	// "encoding/gob"
)

func testConn(i int, x chan bool, address string, port string){
	log.Println("running : ", "---" , i, "---")
	ip_port := address + ":" + port
	c, err := net.Dial("tcp", ip_port)
	if err != nil {
		log.Fatal("net.dial | ", err)
	}
	defer c.Close()
	log.Println("connected to : ", c.RemoteAddr())
	log.Println("we are : " , c.LocalAddr())
	x<-true
}

func client(address *string, port_pstring *string) {
	x := make(chan bool)
	log.Println("Running the App ^^")
	for i := 0; ; i++ {
		go testConn(i, x, *address, *port_pstring)
		<-x
	}
}
