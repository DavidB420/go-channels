package main

import (
	"fmt"
	"time"
)

func main(){

	fmt.Println("Select statement (send to channel) example")

	//Create two channels to use later with the select statement
	channel1 := make(chan int)
	channel2 := make(chan int)

	//Go routine to read from first channel (slower)
	go func(){
		time.Sleep(time.Second * 10)
		value := <-channel1
		fmt.Println(value)
	}()

	//Go routine to read from second channel (faster)
	go func(){
		time.Sleep(time.Second * 5)
		value := <-channel2
		fmt.Println(value)
	}()

	//Select statement that sends values to the two channels, second case should run
	select {
	case channel1 <- 65535:
	case channel2 <- 12345:
	}

}
