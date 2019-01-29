package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listener, e := net.Listen("tcp", "0.0.0.0:8089")
	if e != nil {
		os.Exit(-1)
	}
	for {
		//var res string
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting client: ", err.Error())
			os.Exit(0)
		}
		processor := Processor{
			Conn: conn,
		}
		go processor.process()
		print(conn)
	}
}




