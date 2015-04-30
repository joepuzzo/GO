package main

import (
	"fmt"
	"time"
)

//The main
func main() {
	timeout, _ := time.ParseDuration("5s")
	fmt.Print(int64(timeout))
}
