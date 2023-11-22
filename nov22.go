package main

import (
	"fmt"
	"math"
)

func main() {
	//	fmt.Print("Uclasswork")
	//fmt.Println(func1(2, 4, 6))
	//fmt.Println(func2(3, 0, 4, 0))
	//func3(2, 4, 5)
	func4(8)
}
func func1(a, b, c int) int {
	var katta, kichik = 0, 0
	if a > b && a > c {
		katta = a
	} else if b > c && b > a {
		katta = b
	} else if c > a && c > b {
		katta = c
	}
	if a < b && a < c {
		kichik = a
	} else if b < c && b < a {
		kichik = b
	} else if c < a && c < b {
		kichik = c
	}
	var yigindi int = katta + kichik
	return yigindi
}
func func2(x1, x2, y1, y2 float64) float64 {
	birinchi := math.Pow((x2 - x1), 2)
	ikkinchi := math.Pow((y2 - y1), 2)
	masofa := math.Sqrt(birinchi + ikkinchi)
	return masofa
}
func func3(a, b, c float64) {
	var diskriminant float64 = b*b - 4*a*c
	if diskriminant > 0 {
		var n1 float64 = (-b + math.Sqrt(b)) / (2 * a)
		var n2 float64 = (-b - math.Sqrt(b)) / (2 * a)
		fmt.Println(n1, n2)

	} else if diskriminant == 0 {
		var n3 float64 = -b / 2 * a
		fmt.Println(n3)
	} else if diskriminant < 0 {
		fmt.Println("Yechimga ega emas")
	}
}
func func4(n int) {
	var (
		bir  = 0
		ikki = 1
		uch  = 0
	)
	for i := 1; i < n; i++ {
		uch = bir + ikki
		bir = ikki
		ikki = uch
	}
	fmt.Println(uch, " ")
}
