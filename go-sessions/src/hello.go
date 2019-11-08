package main

import (
	"fmt"
	"os"
	"utils"
)

var x int
var Y string

func multipleReturns(x int, d float64) (s string, i int, b bool, dr float64) {
	s = "test_s"
	i = x
	b = true
	dr = d
	return s, i, b, dr
}

var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

func main() {
	//"_" means discard the value for parameter
	str, _, bv, dv := multipleReturns(1, 1.2)
	fmt.Printf("str = %s bool = %t, dbl = %f\n", str, bv, dv)

	if bv == true {
		fmt.Println("bv is true")
	} else {
		fmt.Println("bv is false")
	}

	//var iPtr *int
	//*iPtr = 10

	//y := iPtr
	//z := y

	//fmt.Println("Z = %d", *iPtr)
	fmt.Println("Hello, 世界")

	utils.Fibtest()

	fmt.Printf("home = %s, user = %s, gopath=%s\n", home, user, gopath)
}
