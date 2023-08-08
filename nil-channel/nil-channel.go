package main

import "fmt"

func main(){

	//Create user input variable
	var userInput int

	fmt.Println("Nil channel Example")

	//Create a nil channel
	var nilChannel chan int

	//Attempt receiving, sending, or closing a nil channel
	fmt.Println("What do you want to attempt\n1.Receiving\n2.Sending\n3.Closing")
	fmt.Scan(&userInput)
	switch userInput{
		case 1:
			<-nilChannel
		case 2:
			nilChannel <- 65535
		case 3:
			close(nilChannel)
		default:
			fmt.Println("Invalid option chosen")
	}
}
