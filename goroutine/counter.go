package main

import (
	"fmt"
	"sync"
	//"time"
)

type Counter struct {
	value int
	lock  sync.Mutex
}

/*Method1: func countValue(c *Counter, ch chan<- bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value++
	fmt.Println(c.value)
	ch <- true
}*/

// Method2
func countValue(c *Counter, wg *sync.WaitGroup) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.value++
	fmt.Println(c.value)
	wg.Done()
}

func main() {
	counter := Counter{}
	//ch := make(chan bool)

	//Mehthod2
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 100; i++ {
		go countValue(&counter, &wg)
	}
	wg.Wait()

	/*Method1:for i := 0; i < 100; i++ {
		go countValue(&counter, ch)
	}
	for i := 0; i < 100; i++ {
		<-ch
	}*/
	//time.Sleep(2 * time.Second)
}
