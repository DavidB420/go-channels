package main

import (
        "fmt"
)

func main(){

        fmt.Println("Buffered Go channel Example")

        //Make a buffered channel
        bufChannel := make (chan int, 2)

        //Send two consecutive values to the buffered channel
	bufChannel <- 65535
	bufChannel <- 12345
	//bufChannel <- 67890 //this will cause deadlock

        //Output the received messages
        fmt.Println(<-bufChannel)
	fmt.Println(<-bufChannel)

}
