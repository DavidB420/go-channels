package main

import "fmt"

func main() {
	fmt.Println("Basic Go Channel Example")
	
	//Create basic Go channel that passes through an integer
	myChannel := make(chan int)

	go func(){
		
		myChannel <- 5 //Send the value 5 to the channel

	}()

	x :=  <- myChannel //Receive the value from the channel

	//Output the received value
	fmt.Println(x)
}
