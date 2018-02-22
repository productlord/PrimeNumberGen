package main

import "fmt"

var i int

func sendtoChannel(ch1 chan<- int) {

	for i = 2; ; i++{

		ch1 <- i

	}
}

func getPrime (fromchan <-chan int, filteredchan chan <-int, primenum int){

for i := range fromchan {

	if i % primenum != 0 {

		filteredchan <- i

	}
}

}

func primenumLister(){

	ch2 := make (chan int)
	
	go sendtoChannel(ch2)
	
	for {
		
		primenum := <-ch2
		fmt.Print(primenum, "\n")
		ch3 := make( chan int)
		go getPrime(ch2,ch3,primenum)
		ch2 = ch3
		if primenum > 50000 {
			break
		}

	}



}

func main(){
	
	primenumLister()
}