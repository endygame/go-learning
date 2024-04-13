package main

import (
	"fmt"
)

func main() {

	var cheezeburger []string;

	var specialSauce []string;

	specialSauce = append(specialSauce, "cheese", "mayo")

	cheezeburger = append(cheezeburger, specialSauce...)

	fmt.Println(cheezeburger)

	var mappy = make(map[string]uint8)

	mappy["julius"] = 34;

	for name, nummy:= range mappy {
		fmt.Printf("%v : %v\n", name, nummy)
	}


	var sliceOfInts = []int32{3,4,8}

	sliceOfInts = append(sliceOfInts, 6, 8, 9)

	fmt.Printf("%v \n", sliceOfInts)


            
}
 