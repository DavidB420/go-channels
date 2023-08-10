package main

import "fmt"

func main(){

	fmt.Println("Channel close Example")

	//Created a buffered channel with a capacity of 2
	bufChannel := make(chan int, 2)

	//Close channel prematurely before sending (will cause panic)
	//close(bufChannel)

	//Send two values to the buffered channel
	bufChannel <- 65535
	bufChannel <- 12345

	//Close the channel
	close(bufChannel)

	//Receive the two values after closing the channel
	fmt.Println(<-bufChannel)
	fmt.Println(<-bufChannel)

	//Indicate that the channel is now closed
	fmt.Println("Channel is now closed.")

}
