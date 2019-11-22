package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	host                           string
	startPort, entPort, timeOutSec, chanCount int
)

func main() {
	fmt.Println("===== START =====")
	flag.StringVar(&host, "h", "127.0.0.1", "scan host address")
	flag.IntVar(&startPort, "sp", 80, "starts scan port")
	flag.IntVar(&entPort, "ep", 9000, "end scan port")
	flag.IntVar(&timeOutSec, "s", 200, "time out second")
	flag.IntVar(&chanCount, "n", 50, "goroutines limit count")

	var (
		ports []int
		wg    sync.WaitGroup
		//ch    chan int
	)
	var ch = make(chan int, chanCount)

	for port := startPort; port < entPort; port++ {
		wg.Add(1)
		ch <- 1
		go func(p int) {
			fmt.Printf("check prot:%d\n", p)
			if checkOpen(host, p, timeOutSec) {
				ports = append(ports, p)
			}
			wg.Done()
			<-ch
		}(port)
	}

	wg.Wait()
	fmt.Println(ports)
	fmt.Println("===== END =====")
}

func checkOpen(h string, p, t int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", h, p), time.Second*2)
	if err != nil {
		return false
	}

	_ = conn.Close()

	return true
}
