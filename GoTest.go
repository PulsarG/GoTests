package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {

	pointMap := map[string]Point{
		"c": {
			X: 4,
			Y: 5,
		},
	}

	pointMap["a"] = Point{
		X: 1,
		Y: 2,
	}
	pointMap["b"] = Point{3, 4}

	fmt.Println(pointMap)

if pointMap == nil {
	pointMap["a"]=Point{1,2}
} else {
	fmt.Println("123")
}

}
