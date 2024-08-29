package main

import "fmt"


func main(){
	randomNum()
	r := newRingBuffer[int](5)
	r.enqueue(1)
	r.enqueue(2)
	r.enqueue(3)
	r.enqueue(4)
	r.enqueue(5)
	r.enqueue(6)
	fmt.Println(r.dequeue())
	fmt.Println(r.dequeue())
	fmt.Println(r.dequeue())
	fmt.Println(r.dequeue())
	fmt.Println(r.dequeue())
}

