// //First Go Program
// //run by doing "go run gointro.go"

// package main

// //Needs double quotes
// import (
// 	"fmt"
// 	"math"
// 	"math/rand"
// )

// //How to call functions from libraries
// func foo() {
// 	fmt.Println("The square root of 4 is ", math.Sqrt(4))
// }

// //Not seeded
// func random() {
// 	fmt.Println("A number from 0-100", rand.Intn(100))
// }

// func add(x, y float64) float64 {
// 	return x + y
// }

// func multiple(a, b string) (string, string) {
// 	return a, b
// }

// func main() {
// 	fmt.Println("Welcome to Go!")
// 	foo()
// 	random()

// 	num1, num2 := 5.6, 9.5
// 	fmt.Println(add(num1, num2))

// 	w1, w2 := "Hey", "there"

// 	fmt.Println(multiple(w1, w2))

// 	// var a int = 62
// 	// var b float64 = float64(a)

// 	// x := a //x will be type int
// }
