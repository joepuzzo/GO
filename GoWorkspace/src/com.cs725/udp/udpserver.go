package main

import (
	//"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	//"strconv"
	"time"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {

	//Define our udp address
	add, err := net.ResolveUDPAddr("udp", ":10001")
	CheckError(err)

	//Listen for connections
	conn, err := net.ListenUDP("udp", add)
	CheckError(err)

	//defer ensures that the connection is closed!
	defer conn.Close()

	//Create a buffer
	var buf []byte = make([]byte, 1500)

	//Loop forever treating connections
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		t2 := time.Now().Nanosecond()
		if err != nil {
			// handle error
		} else {
			fmt.Println("Received ", string(buf[0:n]), " from ", addr)
		}
		//Get the time NOW!
		t3 := time.Now().Nanosecond()
		//Build a json from a map!!
		mapNrm := map[string]int{"t2": t2, "t3": t3}
		mapJsn, _ := json.Marshal(mapNrm)
		//Write out the json!
		_, err = conn.WriteToUDP([]byte(string(mapJsn)), addr)
		CheckError(err)
	}
}
