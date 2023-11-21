package main

import "fmt"

func main() {

	//fmt.Print("jdsksdk")
	//	sayHello()
	//greeting("Davlat")
	// fmt.Println(addNums(23, 5))
	outer()
}
func sayHello() {
	fmt.Println("I am saying hello!")
}
func greeting(name string) {
	fmt.Printf("Hello %s\n", name)
}
func addNums(x int, y int) int {
	sum := x + y
	return sum

}
func generalRandomNumber() int {
	return 0
}
func outer() func(int) int {
	inner := func(n2 int) {
		fmt.Println("some text", n2)
	}
	inner2 := func(n3 int) int {
		fmt.Println("some text 2")
		return n3

	}
	inner(2)

	return inner2
}
