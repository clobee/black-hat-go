package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 1024; i++ {
		wg.Add(1)
		go scanPort(i, &wg)
	}
	wg.Wait()
}

func scanPort(p int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Scan of %d starting... \n", p)

	address := fmt.Sprintf("scanme.nmap.org:%d", p)
	conn, err := net.Dial("tcp", address)

	if err != nil {
		// Port closed / filtered
		return
	}

	conn.Close()
	fmt.Printf("%d is open \n", p)
}
