package main

import (
	"fmt"
	"math"
	"reflect"
)
func main() {
	x:=math.Pow(2,256)
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))

}
