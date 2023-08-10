package main

import "fmt"

func main(){

	fmt.Println("For Loop Channel Example")

	//Create a buffered channel with a capacity of 3
	bufChannel := make(chan int, 3)

	//Send three values to the channel
	bufChannel <- 65535
	bufChannel <- 12345
	bufChannel <- 67890

	//Close the channel
	close(bufChannel) // if we would not close this, for loop will deadlock below
	
	//Iterate through all the items in the channel
	for item := range bufChannel {
		fmt.Println(item)
	}

}
