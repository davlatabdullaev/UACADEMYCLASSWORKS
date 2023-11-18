package main

func main() {
	// 1-masala
	/*
		fmt.Println("Enter a number ")
		var number int
		fmt.Scan(&number)
		if number%3 == 0 && number%3 != 0 {
			fmt.Println("FIZZ")
		}
		if number%5 == 0 && number%5 != 0 {
			fmt.Println("BUZZ")
		}
		if number%3 == 0 && number%5 == 0 {
			fmt.Println("FIZZBUZZ")
		}
	*/

	// 2-masala
	/*
		for true {
			fmt.Println("Yilni kiriting ")
			var number int
			fmt.Scan(&number)
			if (number%4 != 0 || number%100 == 0) && number%400 != 0 {
				fmt.Println("Kabisa yili emas")
			} else if number%4 == 0 || number%400 == 0 {
				fmt.Println("Kabisa yili")
			}
		}
	*/

	// 3-masala
	/*
		    for true{
			fmt.Print("enter a number ")
			var number int
			fmt.Scan(&number)
			if number%2 == 0 {
				fmt.Println(number * number)
			} else {
				fmt.Println(number * number * number)
			}
		}
	*/

	// 4-masala
	/*
		for true {
			fmt.Println("Ismingizni kiriting ")
			var ism string
			fmt.Scan(&ism)
			fmt.Println("Familiyangizni kiriting ")
			var fam string
			fmt.Scan(&fam)
			fmt.Println("Yoshingizni kiriting ")
			var yosh int
			fmt.Scan(&yosh)
			if yosh%2 == 0 {
				fmt.Println(fam)
			} else {
				fmt.Println(ism)
			}
		}
	*/
	// 5-masala
	/*
		for true {
			fmt.Print("Enter a number ")
			var num int
			fmt.Scan(&num)
			if num > 0 {
				fmt.Println(num + 1)
			} else {
				fmt.Println(num)
			}
		}
	*/

	// 6-masala
	/*
		for true {
			fmt.Print("Enter a number ")
			var num int
			fmt.Scan(&num)
			if num > 0 {
				fmt.Println(num + 1)
			} else {
				fmt.Println(num - 2)
			}
		}
	*/

	// 7-masala
	/*
		for true {
			fmt.Print("Enter a number ")
			var num int
			fmt.Scan(&num)
			if num > 0 {
				fmt.Println(num + 1)
			} else if num < 0 {
				fmt.Println(num - 2)
			} else {
				fmt.Println("10")
			}
		}
	*/

	// 8-masala
	/*
		for true {
			fmt.Println("a = ")
			var a int
			fmt.Scan(&a)
			fmt.Println("b = ")
			var b int
			fmt.Scan(&b)
			fmt.Println("c = ")
			var c int
			fmt.Scan(&c)
			if a > 0 && b > 0 && c > 0 {
				fmt.Println(" uchta ")
			} else if (a > 0 && b > 0 && c < 0) || (a > 0 && b < 0 && c > 0) || (a < 0 && b > 0 && c > 0) {
				fmt.Println(" ikkita")
			} else if (a > 0 && b < 0 && c < 0) || (a < 0 && b < 0 && c > 0) || (a < 0 && b > 0 && c < 0) {
				fmt.Println(" bitta ")
			} else if (a < 0 && b < 0 && c < 0) || (a < 0 && b < 0 && c < 0) || (a < 0 && b < 0 && c < 0) {
			fmt.Println("  musbat = 0\n manfiy = 3  ")
		} else if a == 0 || b == 0 || c == 0 {
				fmt.Println(" 0 kiritmang")
			}
		}
	*/

	// 9-masala
	/*
		for true {
			fmt.Println("a = ")
			var a int
			fmt.Scan(&a)
			fmt.Println("b = ")
			var b int
			fmt.Scan(&b)
			fmt.Println("c = ")
			var c int
			fmt.Scan(&c)
			if a > 0 && b > 0 && c > 0 {
				fmt.Println(" musbat = 3\n manfiy = 0 ")
			} else if (a > 0 && b > 0 && c < 0) || (a > 0 && b < 0 && c > 0) || (a < 0 && b > 0 && c > 0) {
				fmt.Println("  musbat = 2\n manfiy = 1 ")
			} else if (a < 0 && b < 0 && c < 0) || (a < 0 && b < 0 && c < 0) || (a < 0 && b < 0 && c < 0) {
				fmt.Println("  musbat = 0\n manfiy = 3  ")
			} else if a == 0 || b == 0 || c == 0 {
				fmt.Println(" 0 kiritmang")
			}
		}
	*/

	// 10-masala
	/*
		for true {
			fmt.Println(" a = ")
			var a int
			fmt.Scan(&a)
			fmt.Println(" b = ")
			var b int
			fmt.Scan(&b)
			if a > b {
				fmt.Println(a)
			} else if a < b {
				fmt.Println(b)
			}
		}
	*/

	// 11-masala
	/*
		for true {
			fmt.Println("Birinchi sonni kiriting")
			var a int
			fmt.Scan(&a)
			fmt.Println("Ikkinchi sonni kiriting")
			var b int
			fmt.Scan(&b)
			if a > b {
				fmt.Println("Ikkinchi")
			} else if a < b {
				fmt.Println("Birinchi")
			}
		}
	*/

	// 12-masala
	/*
		for true {
			fmt.Println("Birinchi sonni kiriting")
			var a int
			fmt.Scan(&a)
			fmt.Println("Ikkinchi sonni kiriting")
			var b int
			fmt.Scan(&b)
			if a > b {
				fmt.Println(a, b)
			} else if a < b {
				fmt.Println(b, a)
			}
		}
	*/

	// FOR

	// 1-masala
	/*
		fmt.Println(" k = ")
		var k int
		fmt.Scan(&k)
		fmt.Println(" n = ")
		var n int
		fmt.Scan(&n)

		for i := 0; i < n; i++ {
			fmt.Println(k)
		}
	*/

	// 2-masala
	/*
			for true {
			fmt.Println(" a = ")
			var a int
			fmt.Scan(&a)
			fmt.Println(" b = ")
			var b int
			fmt.Scan(&b)
		    if(a<b){
			for i := a; i <= b; i++ {
				fmt.Println(i)
			}}}
	*/

	// 3-masala
	/*
		for true {
			fmt.Println(" a = ")
			var a int
			fmt.Scan(&a)
			fmt.Println(" b = ")
			var b int
			fmt.Scan(&b)
			if a < b {
				for i := b - 1; i > a; i-- {
					fmt.Println(i)
				}
			}
		}
	*/

	// 4-masala
	/*
		for true {
			fmt.Print("Bir kg konfet narxini kiriting ")
			var n int
			fmt.Scan(&n)
			for i := 1; i <= 10; i++ {
				fmt.Println(i, i*n)
			}
		}
	*/

	// 5-masala
	/*
		for true {
			fmt.Print("Bir kg konfet narxini kiriting ")
			var n int
			fmt.Scan(&n)
			for i := 0.1; i <= 1.0; i += 0.1 {

				fmt.Printf("%.1f kg konfetning narxi %.1f so'm \n", i, i*float64(n))

			}
		}
	*/
	// SWITCH
	// 1-masala
	/*
		for true {
			fmt.Println("Hafta kunini kiriting ")
			var n int
			fmt.Scan(&n)
			switch n {
			case 1:
				fmt.Println("DUSHANBA")
			case 2:
				fmt.Println("SESHANBA")
			case 3:
				fmt.Println("CHORSHANBA")
			case 4:
				fmt.Println("PAYSHANBA")
			case 5:
				fmt.Println("JUMA")
			case 6:
				fmt.Println("SHANBA")
			case 7:
				fmt.Println("YAKSHANBA")
			default:
				fmt.Println("XATO")
			}
		}
	*/

	// 2-masala
	/*
		for true {
			fmt.Println("Bahoni kiriting ")
			var n int
			fmt.Scan(&n)
			switch n {
			case 1:
				fmt.Println("YOMON")
			case 2:
				fmt.Println("QONIQARSIZ")
			case 3:
				fmt.Println("QONIQARLI")
			case 4:
				fmt.Println("YAXSHI")
			case 5:
				fmt.Println("A'LO")
			default:
				fmt.Println("XATO")
			}
		}
	*/

	// 3-masala
	/*
		for true {
			fmt.Println("Oyni kiriting ")
			var n int
			fmt.Scan(&n)
			switch n {
			case 1:
				fmt.Println("Yanvar - QISH")
			case 2:
				fmt.Println("Fevral - QISH")
			case 3:
				fmt.Println("Mart - BAHOR")
			case 4:
				fmt.Println("Aprel - BAHOR")
			case 5:
				fmt.Println("May - BAHOR")
			case 6:
				fmt.Println("Iyun - YOZ")
			case 7:
				fmt.Println("Iyul - YOZ")
			case 8:
				fmt.Println("Avgust - YOZ")
			case 9:
				fmt.Println("Sentabr - KUZ")
			case 10:
				fmt.Println("Oktabr - KUZ")
			case 11:
				fmt.Println("Noyabr - KUZ")
			case 12:
				fmt.Println("Dekabr - QISH")
			default:
				fmt.Println("XATO")
			}

		} */

	// 4-masala
	/*
		for true {
			fmt.Println("Oyni kiriting ")
			var n int
			fmt.Scan(&n)
			switch n {
			case 1:
				fmt.Println("Yanvar - 31 kun bor")
			case 2:
				fmt.Println("Fevral - 28 yoki 29 kun bor")
			case 3:
				fmt.Println("Mart - 31 kun bor")
			case 4:
				fmt.Println("Aprel - 30 kun bor")
			case 5:
				fmt.Println("May - 31 kun bor")
			case 6:
				fmt.Println("Iyun - 30 kun bor")
			case 7:
				fmt.Println("Iyul - 31 kun bor")
			case 8:
				fmt.Println("Avgust - 31 kun bor")
			case 9:
				fmt.Println("Sentabr - 30 kun bor")
			case 10:
				fmt.Println("Oktabr - 31 kun bor")
			case 11:
				fmt.Println("Noyabr - 30 kun bor")
			case 12:
				fmt.Println("Dekabr - 31 kun bor")
			default:
				fmt.Println("XATO")
			}

		}*/

	// 5-masala
	/*
		for true {
			fmt.Println("a sonni kiriting ")
			var a int
			fmt.Scan(&a)
			fmt.Println("b sonni kiriting ")
			var b int
			fmt.Scan(&b)
			fmt.Println("Amalni kiriting: \n 1-qo'shish \n 2-ayirish \n 3-bo'lish \n 4-ko'paytirish ")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				fmt.Println("Yig'indi ", a+b)
			case 2:
				fmt.Println("Ayirma ", a-b)
			case 3:
				fmt.Println("Bo'linma ", a/b)
			case 4:
				fmt.Println("Ko'paytma ", a*b)
			default:
				fmt.Println("Bunday amal kiritilmadi")

			}
		}
	*/
}
