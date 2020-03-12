package main

import (
	"fmt"
	"net"
)

func main() {
	
		address := fmt.Sprintf("scanme.nmap.org:%d", 80)

		conn, err := net.Dial("tcp", address)
		if err == nil {
			fmt.Printf("%d open\n", 80)
		}
		conn.Close()
		
	
}
