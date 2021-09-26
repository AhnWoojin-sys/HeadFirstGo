package main

import (
	"fmt"
	"log"
)

func find(item string, slice []string) bool {
	for _, sliceitem := range slice {
		if sliceitem == item {
			return true
		}
	}
	return false
}

type Refrigerator []string

func (r Refrigerator) Open() {
	fmt.Println("Opening refrigerator")
}

func (r Refrigerator) Close() {
	fmt.Println("Closing refrigerator")
}

func (r Refrigerator) FindFood(food string) error {
	r.Open()
	defer r.Close()
	if find(food, r){
		fmt.Println("Found", food);
	} else {
		return fmt.Errorf("%s not found", food)
	}
	return nil
}

func main() {
	fridge := Refrigerator{"Milk", "Pizza", "Salsa"}
	for _, food := range []string{"Milk", "Pizza"}{
		err := fridge.FindFood(food)
		if err != nil {
			log.Fatal(err)
		}
	}
}