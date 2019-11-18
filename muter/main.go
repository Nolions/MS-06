package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex

	fmt.Println("Lock the lock. (G0)")
	m.Lock()
	fmt.Println("The lock is locked.(G0)")

	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock. (G%d)\n", i)
			m.Lock()
			fmt.Printf("The lock is locked. (G%d)\n", i)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlock the lock. (G0)")
	m.Unlock()
	fmt.Println("The lock is unlocked. (G0)")
	time.Sleep(time.Second)
}
