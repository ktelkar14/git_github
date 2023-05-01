package main

import "fmt"

func main() {

	/*var total float32 = 0
	//var x [5]float32
	//{98,34,56,45,37}

	x := [5]float32{98, 93, 77, 82, 83}
	for _, value := range x {
		total += value
	}
	fmt.Println("Total value ", total/float32(len(x)))
	*/

	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
	y := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(y[2:5])

}
