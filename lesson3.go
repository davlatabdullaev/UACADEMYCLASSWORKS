package main

func main() {
	// 1-masala
	/*  var (
	      a int
	      b int
	    )
	    fmt.Println("A ni kiriting ")
	    fmt.Scan(&a)
	    fmt.Println("B ni kiriting ")
	    fmt.Scan(&b)

	    a, b = b, a

	    fmt.Println("A = ", a)
	    fmt.Println("B = ", b) */

	// 2-masala
	/*    var (
	        a int
	        b int
	        c int
	      )
	      fmt.Println("A ni kiriting ")
	      fmt.Scan(&a)
	      fmt.Println("B ni kiriting ")
	      fmt.Scan(&b)
	      fmt.Println("C ni kiriting ")
	      fmt.Scan(&c)

	      a, b, c = c, a, b

	      fmt.Println("A = ", a)
	      fmt.Println("B = ", b)
	      fmt.Println("C = ", c) */

	// 3-masala
	/*  fmt.Println("Ikki xonali son kiriting")
	    var (
	      son int
	    )
	    fmt.Scan(&son)
	    fmt.Print("onliklar xonasidagi raqam ", son/10)
	    fmt.Print("birliklar xonasidagi raqam ", son%10)
	*/

	// 4-masala
	/*  fmt.Println("Ikki xonali son kiriting")
	    var (
	      son int
	    )
	    fmt.Scan(&son)
	    fmt.Println("raqamlar yig'indisi ", son%10+son/10)*/

	// 5-masala
	/*  fmt.Println("Ikki xonali son kiriting")
	    var (
	      son int
	    )
	    fmt.Scan(&son)
	    var natija int = (son%10)*10 + son/10
	    fmt.Println("teskari son ", natija) */

	// 6-masala
	/*  fmt.Println("Uch xonali son kiriting")
	    var (
	      son int
	    )
	    fmt.Scan(&son)
	    fmt.Print("Yuzlar xonasidagi raqam ", son/100)
	*/

	// 7-masala
	/*  fmt.Println("Uch xonali son kiriting")
	    var (
	      son int
	    )
	    fmt.Scan(&son)
	    fmt.Print("Onliklar xonasidagi raqam ", (son/10)%10)
	    fmt.Print("Birliklar xonasidagi raqam ", (son%100)%10) */

	// 8-masala
	/* fmt.Println("Uch xonali son kiriting")
	var (
	  son int
	)
	fmt.Scan(&son)
	fmt.Print("Raqamlar yig'indisi", (son/10)%10+(son%100)%10)+son/100)
	*/

	// 9-masala
	/*  fmt.Println("Uch xonali son kiriting")
	    var (
	      son int
	    )
	    fmt.Scan(&son)
	    var natija = (son%100)%10*100 + (son%100)/10*10 + (son / 100)
	    fmt.Println("Teskari son ", (natija)) */

	//10-masala
	/*  fmt.Println("Uch xonali son kiriting")
	    var (
	      son int
	    )
	    fmt.Scan(&son)
	    var yangison int = (son%100)/10*100 + (son%100)%10*10 + son/100
	    fmt.Print("Yangi son ", yangison) */

	//11-masala
	/*  fmt.Println("Uch xonali son kiriting")
	    var (
	      son int
	    )
	    fmt.Scan(&son)
	    var yangison2 int = (son%100)%10*100 + son/100*10 + (son%100)/10
	    fmt.Println("Yangi son ", yangison2) */

	//12-masala
	/*  fmt.Println("Uch xonali son kiriting")
	    var (
	      son int
	    )
	    fmt.Scan(&son)
	    var yangison3 int = (son%100)/10*100 + (son/100)*10 + son%10
	    fmt.Println("Yangi son ", yangison3) */
}
