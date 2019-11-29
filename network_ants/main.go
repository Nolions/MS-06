package main

import (
	"flag"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"net"
	"sync"
	"time"
)

var (
	host                                      string
	startPort, entPort, timeOutSec, chanCount int
	ports                                     []int
)

func main() {
	fmt.Println("===== START =====")
	inputParam()

	var wg sync.WaitGroup

	// Use the pool with a function,
	p, _ := ants.NewPoolWithFunc(chanCount, func(i interface{}) {

		checkOpen(host, i.(int), timeOutSec)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for i := startPort; i < entPort; i++ {
		wg.Add(1)
		_ = p.Invoke(int(i))
	}

	wg.Wait()
	p.Running()

	fmt.Println(ports)
	fmt.Println("===== END =====")
}

func inputParam() {
	flag.StringVar(&host, "h", "127.0.0.1", "scan host address")
	flag.IntVar(&startPort, "sp", 80, "starts scan port")
	flag.IntVar(&entPort, "ep", 9000, "end scan port")
	flag.IntVar(&timeOutSec, "s", 200, "time out second")
	flag.IntVar(&chanCount, "n", 50, "goroutines limit count")
}

func checkOpen(h string, p, t int) bool {
	fmt.Printf("check prot:%d\n", p)

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", h, p), time.Second*2)
	if err != nil {
		return false
	}
	_ = conn.Close()

	ports = append(ports, p)
	return true
}
