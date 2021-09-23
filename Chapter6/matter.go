package main

import "fmt"

func main(){
	array := [5]string{"a", "b", "c", "d", "e"}
	slice := array[1:3] // slice {b, c}
	slice = append(slice, "x") // slice {b, c, x}
	slice = append(slice, "y", "z") // slice {"b, c, x, y, z}
	for _, letter := range slice {
		fmt.Println(letter)
	}
}

/* 출력값
*/ 