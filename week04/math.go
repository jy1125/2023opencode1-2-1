package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	var a int
	var b float32
	var c, d string
	var f string
	d = "Go..."
	c = "inha"
	b = 2.7
	a = 9
	fmt.Println(a, b, c, d, f)
	fmt.Println(a > int(b))
	fmt.Println(reflect.TypeOf('B'))
	fmt.Println(reflect.TypeOf(100))
	fmt.Println(reflect.TypeOf(2.71))
	fmt.Println(reflect.TypeOf(false))
	fmt.Println(math.Floor(2.75))
	fmt.Println(math.Ceil(2.75))
	fmt.Println(strings.Title("오픈소스프로그래밍~"))
}
