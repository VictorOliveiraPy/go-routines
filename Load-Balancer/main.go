package main

import (
	"fmt"
	"time"
)


func worker(workerID int, data chan int){
	for x := range data {  // processa os dados
		fmt.Printf("Worker %d  received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {
	channel := make(chan int)

	WtdWorkers := 1000

	// Inicializa os workers
	for i := 0; i < WtdWorkers; i++ {
        go worker(i, channel)
    }

	for i := 0; i < 1000000; i++ {  // Enviar informacoes para o canal
		channel <- i   //recebendo informaces
    }

}