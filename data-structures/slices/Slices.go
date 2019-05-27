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

	slice5 := make([]int, 10)
	slice5[9] = 12
	fmt.Println(slice5)

	//Create a slice with size of 10 and a array with size of 20
	//When we create a slice, internally a internal array will be created
	slice5 = make([]int, 10, 20)
	fmt.Println(slice5, len(slice5), cap(slice5))

	slice5 = append(slice5, 2, 4, 6, 7, 2, 8, 9, 2, 67, 78)
	fmt.Println(slice5, len(slice5), cap(slice5))

	//Slices can dynamically increses the size of the internal array
	slice5 = append(slice5, 34)
	fmt.Println(slice5, len(slice5), cap(slice5))
}
