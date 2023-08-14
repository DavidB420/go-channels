package main

import (
	"fmt"
	"time"
)

func main(){
	
	fmt.Println("Select statement (receive from channel) example")

	//Create two channels to use later with the select statement
	channel1 := make(chan int)
	channel2 := make(chan int)

	//Go routine to send to first channel (faster)
	go func(){
		time.Sleep(time.Second * 5)
		channel1 <- 65535
	}()

	//Go routine to send to second channel (slower)
	go func(){
		time.Sleep(time.Second * 10)
		channel2 <- 12345
	}()

	//Closing any of the channels will cause select to return 0
	//close(channel1)
	//close(channel2)
	//Select statement that receives values from the two channels, first case should run
	select {
		case value := <-channel1:
			fmt.Println(value)
		case value := <-channel2:
			fmt.Println(value)
	}

	//Give time for one of the goroutines to finish
	time.Sleep(time.Second * 1)

}
