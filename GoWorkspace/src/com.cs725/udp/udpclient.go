package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	var address string
	fmt.Printf("Please enter an address or hostname: ")
	fmt.Scanf("%s", &address)
	Connect2Server(address, "10001")
}

func Connect2Server(srv_addr string, port string) {

	//Build a UDP Address for the server
	saddr, err := net.ResolveUDPAddr("udp", srv_addr+":"+port)
	CheckError(err)

	//Create a udp connection "object"
	conn, err := net.DialUDP("udp", nil, saddr)
	CheckError(err)

	//This makes sure this gets closed no matter what
	defer conn.Close()

	for {
		//Random shit to write out
		t1 := time.Now().Nanosecond()
		buf1 := []byte(strconv.Itoa(t1))
		buf2 := make([]byte, 1500)

		//Write a value out the the server
		_, err := conn.Write(buf1)

		if err != nil {
			//If we get an error then print it
			fmt.Println(t1, err)
			failed2connect()
		}

		//Get the data back from the server
		n, _, err := conn.ReadFromUDP(buf2)
		t4 := time.Now().Nanosecond()
		if err != nil {
			// handle error
		} else {
			t2, t3 := getData(string(buf2[0:n]))
			timeandoffset(t1, t2, t3, t4)
		}
		time.Sleep(time.Second * 5)
	}
}

/**
* This function will parse the given string as a json string
 */
func getData(str string) (t2, t3 int) {
	fmt.Println(str)
	type Message struct {
		T2 int
		T3 int
	}

	m := Message{}

	json.Unmarshal([]byte(str), &m)
	fmt.Println(m)
	fmt.Printf("t2: %d, t3: %d\n", m.T2, m.T3)
	t2 = m.T2
	t3 = m.T3
	return
}

/**
* round trip time (RTT) as (t4 - t1) - (t3 - t2)
* clock offset as (t1 + RTT/2) - t2:
 */
func timeandoffset(t1, t2, t3, t4 int) {
	rtt := (t4 - t1) - (t3 - t2)
	off := (t1 + rtt/2) - t2
	fmt.Printf("RTT = %d ClockOffset = %d\n", rtt, off)
}

func failed2connect() {
	var ans string
	fmt.Println("Failed to connect! Woule You like to try again? type \"yes\" or \"no\" ")
	fmt.Scanf("%s", &ans)
	if ans == "no" {
		os.Exit(0)
	}
}
