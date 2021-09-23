package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers {
		fmt.Println(number);
		sum += number
	}

	fmt.Printf("Average : %0.2f", sum/float64(len(numbers)));
}