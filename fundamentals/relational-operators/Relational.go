package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings", "Go==Go", "go" == "go")
	fmt.Println("!=", "3!=2", 3 != 2)
	fmt.Println("<", "3<2", 3 < 2)
	fmt.Println(">", "3>2", 3 > 2)
	fmt.Println("<=", "3<=2", 3 <= 2)
	fmt.Println(">=", "3>=2", 3 >= 2)

	d1 := time.Unix(0, 0) //1969-12-31 21:00:00 -0300 -03
	d2 := time.Unix(0, 0)

	fmt.Println("Date", fmt.Sprint(d1)+"=="+fmt.Sprint(d2), d1 == d2)
	fmt.Println("Date", fmt.Sprint(d1)+".Equal("+fmt.Sprint(d2)+")", d1.Equal(d2))

	type Entity struct {
		value string
	}

	ea := Entity{"IDE"}
	eb := Entity{"IDE"}

	fmt.Println("Struct", fmt.Sprint(ea)+"=="+fmt.Sprint(eb), ea == eb)
}
