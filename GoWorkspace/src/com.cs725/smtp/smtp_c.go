package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

type SMTPMessage struct {
	Type    string
	Message string
}

//The main
func main() {
	var address string
	var port int
	fmt.Printf("Please enter an address or hostname: ")
	fmt.Scanf("%s", &address)
	fmt.Printf("Please enter a port number: ")
	fmt.Scanf("%d", &port)
	ConnectToServer(address, port)
}

func ConnectToServer(server_address string, port int) {
	//Connects to the srever
	conn, err := net.Dial("tcp", server_address+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	//Print some basic data
	fmt.Println("Local Address: " + conn.LocalAddr().String())
	fmt.Println("Remote Address: ", conn.RemoteAddr())

	//Create a buffer reader from the connection
	r := bufio.NewReader(conn)
	read(r)

	//Write a get message to the connection buffer retured
	writeAndGet(r, conn, "HELO jgb53.cs")
	writeAndGet(r, conn, "MAIL FROM: <helloworld@unh.com>")
	writeAndGet(r, conn, "RCPT TO: <jgb53@wildcats.unh.edu>")
	writeAndGet(r, conn, "DATA")
	write(conn, "WHATS UP DOG!")
	write(conn, "HOW U BEEN??")
	write(conn, "TESTING NEWLINE \n YEAH!")
	write(conn, "\n Regards, \n -Joe")
	writeAndGet(r, conn, "\n.")
	writeAndGet(r, conn, "QUIT")
}

func checkMessage(stat string, err error) (ok bool) {
	if err != nil {
		fmt.Println("Error: ", err)
		ok = false
	} else {
		fmt.Println("Recieved: " + stat)
		ok = true
	}
	return
}

func writeAndGet(in *bufio.Reader, conn net.Conn, msg string) {
	write(conn, msg)
	read(in)
}

func read(in *bufio.Reader) {
	status, err := in.ReadString('\n')
	checkMessage(status, err)
}

func write(conn net.Conn, msg string) {
	fmt.Fprintf(conn, msg+"\n\r")
	fmt.Println("SENT: " + msg)
}
