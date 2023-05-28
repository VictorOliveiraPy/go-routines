package main

import (
	"fmt"
	"sync"
)

// go routine 1
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go publish(ch)
	reader(ch, &wg)
	wg.Wait()
	
	}



func reader(ch chan int, wg *sync.WaitGroup){
	for x := range ch {
		fmt.Println("Received %d", x)  // ler o canal pra esvazia
		wg.Done()
	}
}

func publish(ch chan int){
	for i := 0; i <10; i++ { // so publica quando o canal tiver vazio
		ch <- i
	}
	close(ch)
}