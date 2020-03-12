package main

import (
	"fmt"
	"sync"
	"net"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err == nil {
			
			fmt.Printf("%d open\n",p)
		}
		conn.Close()
		
	}
	wg.Done()
}

func main() {
	ports := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}
