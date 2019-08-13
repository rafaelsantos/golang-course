package main

import (
	"fmt"
	"runtime"
)

func main() {
	//Show total of processors
	fmt.Println(runtime.NumCPU())
}
