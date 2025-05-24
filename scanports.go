package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		address := fmt.Sprintf("127.0.0.1:%d", p)
		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		if err == nil {
			fmt.Printf("Port %d open\n", p)
			conn.Close()
		}
		wg.Done()
	}
}

func main() {
	ports := make(chan int, 100)
	var wg sync.WaitGroup

	// Khởi tạo 100 worker goroutines
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}

	// Gửi cổng vào channel và tăng wg
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}

	// Chờ hoàn thành
	wg.Wait()
	close(ports)
}
