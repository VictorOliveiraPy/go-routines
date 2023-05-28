package main

func main() {
	channel := make(chan string, 2)
	channel <- "hello world"
	channel <- "golang.org"

	println(<-channel)
	println(<- channel)

}