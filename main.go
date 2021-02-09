package main

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/gookit/color"
)

func main() {
	// clear screen
	fmt.Print("\033[H\033[2J")

	var IP string
	// PI = Port Input
	var PI string

	color.Bold.Print("TARGET IP> ")
	fmt.Scan(&IP)

	color.Bold.Print("PORT(s)> ")
	fmt.Scan(&PI)

	PORTS := strings.Split(PI, ",")

	for _, PORT := range PORTS {

		TARGET := fmt.Sprintf("%s:%s", IP, PORT)

		// // Change the value to 10 if you want to use it on your local network.
		_, err := net.DialTimeout("tcp", TARGET, 600*time.Millisecond)

		if err != nil {
			color.Red.Printf("Port %s is closed\n", PORT)
		} else {
			color.Green.Printf("Port %s is open\n", PORT)
		}
	}
}
