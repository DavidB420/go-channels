package main

import "fmt"

func main(){
	
	fmt.Println("Go Unbuffered Channel Deadlock Example")

	//Create unbuffered channel
	unbufChannel := make(chan int)

	//Send value to channel on main goroutine
	unbufChannel <- 65535

	//Receive value should cause deadlock
	fmt.Println(<-unbufChannel)

}
