package main

import "testing"


func TestsendtoChannel(ch1 chan<- int, t *testing.T) {
var v int

	for i := 2; ; i++{

		ch1 <- i

		ch4 := make (chan int)

		TestsendtoChannel(ch4, t)

		v = <-ch4

		if v == 0 {t.Error("Error Testing Works!")}
		
	}

	
}






