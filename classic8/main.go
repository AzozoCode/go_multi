package main

import "fmt"

type Coordinates struct {
	Latitude, Longitude float64
}

func main() {
	//create a map with struct key

	location := make(map[Coordinates]string)

	//add location to the map

	location[Coordinates{40.2345, 32.9045}] = "New York"
	location[Coordinates{79.00, 67.095}] = "Newark"

	for key, value := range location {
		fmt.Printf("key:%v value:%s\n", key, value)
	}
}
