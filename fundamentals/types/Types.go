package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//Integer numbers
	fmt.Println(32, 56, 567)
	fmt.Println("Integer literal is", reflect.TypeOf(56789))

	//Unsigned Integer 8 bits (uint8)
	var b byte = 255
	fmt.Println("Byte is", reflect.TypeOf(b))

	//Unsigned Integer 16 bits (uint16)
	var x uint16 = math.MaxUint16
	fmt.Println("Value of x is", x)
	fmt.Println("Type of x is", reflect.TypeOf(x))

	//Unsigned Integer 32 bits (uint32)
	var y uint32 = math.MaxUint32
	fmt.Println("Value of y is", y)
	fmt.Println("Type of y is", reflect.TypeOf(y))

	//Unsigned Integer 64 bits (uint64)
	var z uint64 = math.MaxUint64
	fmt.Println("Value of z is", z)
	fmt.Println("Type of z is", reflect.TypeOf(z))

	//Signed Integer 8 bits (int8)
	var i int8 = 2
	fmt.Println("Value of i is", i)
	fmt.Println("Type is", reflect.TypeOf(i))

	//Signed Integer 16 bits (int16)
	var q int16 = math.MaxInt16
	fmt.Println("Value of q is", q)
	fmt.Println("Type of q is", reflect.TypeOf(q))

	//Signed Integer 32 bits (int32)
	var r int32 = math.MaxInt32
	fmt.Println("Value of r is", r)
	fmt.Println("Type of r is", reflect.TypeOf(r))

	//Signed Integer 64 bits (int64)
	var k int64 = math.MaxInt64
	fmt.Println("Value of k is", k)
	fmt.Println("Type of k is", reflect.TypeOf(k))

	//Mapping in unicode table of type int32
	var s rune
	s = 't'
	fmt.Println("Value of s is", s)
	fmt.Println("Type of s", reflect.TypeOf(s))

	var o float32 = 56.78
	fmt.Println("Type of o is", reflect.TypeOf(o))

	//Default type of literal float values is float64
	fmt.Println("Type of literal value 56.78 is", reflect.TypeOf(56.78))

	//Boolean
	valid := true
	fmt.Println("Type of valid is", reflect.TypeOf(valid))
	fmt.Println(!valid)

	//String
	text := "Golang"
	fmt.Println(text+"!", "!")
	fmt.Println(len(text), text)

	//String with multiple lines
	phrase := `This text
			   with multiple lines
			   !!!!!!!!!!!!`
	fmt.Println(phrase)
}
