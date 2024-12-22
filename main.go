package main

import "fmt"

func main() {
	const pi float64 = 3.1415
	var radius float64 = 5
	var area float64 = (radius * radius * pi)

	fmt.Println("pi = ", pi, ", radius = ", radius, ", area = ", area)

}
