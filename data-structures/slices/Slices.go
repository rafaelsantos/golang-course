package main

import (
	"fmt"
	"reflect"
)

func main() {
	array := [3]int{1, 2, 3} //array
	slice := []int{1, 2, 3}  //slice (dynamic size)

	fmt.Println(array, slice)
	fmt.Println(reflect.TypeOf(array), reflect.TypeOf(slice))

	array2 := [5]int{1, 3, 6, 7, 8}
	slice2 := array2[1:3] //Slice defines a part of array
	fmt.Println(array2, slice2)

	//Get from init to index 2
	slice3 := array2[:2]
	fmt.Println(array2, slice3)

	slice4 := slice2[:1]
	fmt.Println(slice2, slice4)
}
