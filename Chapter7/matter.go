package main

import (
	"fmt"
)

func main(){
	jewelry := make(map[string]float64)
	jewelry["necklace"] = 89.99
	jewelry["earrings"] = 79.99
	clothing := map[string]float64{"Pants": 59.99, "Shirts": 39.99}

	for i, jewel := jewelry range {
		fmt.Println(i, jewel);
	}

	fmt.Println(jewelry["necklace"])
	fmt.Println(jewelry["earring"])
	fmt.Println(clothing["Shirts"])
	fmt.Println(clothing["Pants"])
}