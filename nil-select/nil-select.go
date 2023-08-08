package main

import (
	"fmt"
	"time"
)

func main(){
	
	fmt.Println("Nil Switch Example")

	//Create two channels that will be used in the select statement
	channel1 := make(chan int)
	channel2 := make(chan int)

	//Go routine to send value to first channel
	go func(){
		channel1 <- 65535
	}()

	//Go routine to send value to second channel
	go func(){
		channel2 <- 12345
	}()

	//Select statement that utilizes nil channels to stop the program
	for channel1 != nil || channel2 != nil{
		select{
		case value := <-channel1:
			fmt.Println(value)
			channel1 = nil
		case value := <-channel2:
			fmt.Println(value)
			channel2 = nil
		}
	}

	//Give time for one of the goroutines to finish
	time.Sleep(time.Second * 1)

}
