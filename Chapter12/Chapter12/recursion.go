package main

import "fmt"

func recurses(count int) {
	fmt.Println("Oh, no, I`m stuck!")
	if count != 0 {
		recurses(count - 1)
	}
}

func main() {
	recurses(3)
}