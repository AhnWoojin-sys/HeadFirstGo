package main

import (
	"fmt"
	"geo"
	"log"
)

func main() {
	coordinates := geo.Coordinates{}
	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}

	err = coordinates.SetLongitude(22.08)
	if err!= nil {
		log.Fatal(err)
	}

	fmt.Println(coordinates.Latitude());
	fmt.Println(coordinates.Longitude());
}