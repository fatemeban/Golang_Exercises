package main

import (
	"fmt"
	"math/rand"
)

func randomNum(){
	var ints []int
	//ints := []int{}
	//ints := make([]int,100)
	//ints := make([]int,0,100)

	for  i := 0 ; i < 100; i++{
		ints = append(ints, rand.Intn(101))
		//ints[i] = rand.Intn(101)
	}

	for _, value := range ints {
		if value%2 == 0 && value%3 == 0 {
			fmt.Println("Six!")
		} else if value%2 == 0 {
			fmt.Println("Two!")
		} else if value%3 == 0 {
			fmt.Println("Three!")
		} else {
			fmt.Println("Never mind")
		}
	}


	fmt.Println(ints)
}