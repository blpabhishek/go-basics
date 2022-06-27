package main

import "fmt"
import "example/tw/basic"

func main() {
	fmt.Println(basic.IsEven(4))
	fmt.Println(basic.Square(2))
	fmt.Println(basic.Cube(2))
	fmt.Println(basic.GCD(12, 12))
	fmt.Println(basic.LCM(4, 3))
	fmt.Println(basic.LCM(3, 7))
	fmt.Println(basic.LCM(28, 16))
	fmt.Println(basic.Sum(10))
	fmt.Println(basic.TakeFibonacci(10))
}
