package main

import "fmt"

func snack() {
	defer fmt.Println("Closing refrigerator")
	fmt.Println("Opening refrigerator")
	panic("refrigerator is empty")
}

func main() {
	snack()
}

/*
Output : 
Opening refrigerator
Closing refrigerator
panicL: refrigerator is empty

goroutine 1 [running]:
main.snack()

*/