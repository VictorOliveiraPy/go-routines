package main

import "fmt"

// go routine 1
func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch)
	
	}



func reader(ch chan int){
	for x := range ch {
		fmt.Println("Received %d", x)  // ler o canal pra esvazia
	}
}

func publish(ch chan int){
	for i := 0; 1 <10; i++ { // so publica quando o canal tiver vazio
		ch <- i
	}
	close(ch)
}