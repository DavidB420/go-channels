package main

import "fmt"

func main(){
	
	fmt.Println("For loop channel Example (No Range)")

	//Create a buffered channel with a capacity of 3
	bufChannel := make(chan int, 3)

	//Send three values to the channel
	for i := 0; i < 3; i++ {
		bufChannel <- i
	}

	//Receive three values from the channel
	for i := 0; i < 3; i++ {
		fmt.Println(<-bufChannel)
	}

}
