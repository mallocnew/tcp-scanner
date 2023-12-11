package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("127.0.0.1:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%s closed\n", address)
				return
			}
			conn.Close()
			fmt.Printf("%s opened\n", address)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start) / 1e6
	fmt.Printf("\n\n%d ms\n", elapsed)
}
