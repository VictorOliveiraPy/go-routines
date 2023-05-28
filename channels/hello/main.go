package main

import "fmt"

// go routine 1
func main() {
	channel := make(chan string)

	// go routine 2
	go func(){
		channel <- "Ola Mundo!" // ESTA CHEIO
	}()

	// go routine 1
	msg := <-channel // canal esvazia
	fmt.Println(msg)
}