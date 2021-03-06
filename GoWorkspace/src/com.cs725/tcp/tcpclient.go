package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

//The main
func main() {
	var address string
	fmt.Printf("Please enter an address or hostname: ")
	fmt.Scanf("%s", &address)
	ConnectToServer(address, 80)
}

func ConnectToServer(server_address string, port int) {
	//Connects to the srever
	conn, err := net.Dial("tcp", server_address+":"+strconv.Itoa(port))
	if err != nil {
		//Handle error
		fmt.Println("Error:", err)
		return
	}

	//Print some basic data
	fmt.Println("Local Address: " + conn.LocalAddr().String())
	fmt.Println("Remote Address: ", conn.RemoteAddr())

	//Write a get message to the connection buffer retured
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')

	fmt.Println(status)
	fmt.Println(err)
}
