package main

import (
	"fmt"
	"time"
)


func task(name string){
	for i := 0; i < 10; i++ {
		fmt.Printf("%d : Task %s is running", i, name)
		time.Sleep(1 * time.Second)
	}
}

//Thread 1
func main() {
	//Thread 2
   go task("A")

   //Thread 3
   go task("B")
   //  nADA AQUI.
   // SAIR
   time.Sleep(5 * time.Second)
}