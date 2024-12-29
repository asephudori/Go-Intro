// Go if statement
package main
import ("fmt")

func main()  {
	if 20 > 18 {
		fmt.Println("20 lebih besar dari 18")
	}
}

package main
import ("fmt")

func main()  {
	x:= 20
	y:= 18

	if x > y {
		fmt.Println("x lebih besar dari y")
	}
}

// Go if else Statement
package main
import ("fmt")

func main()  {
	time:= 20

	if (time < 18) {
		fmt.Println("otsukare")
	} else{
		fmt.Println("oyasumi")
	}
}

// Go else if Statement
package main
import ("fmt")

func main()  {
	time:= 22

	if (time < 10) {
		fmt.Println("ohayou")
	} else if time < 20{
		fmt.Println("otsukare")
	} else{
		fmt.Println("oyasumi")			
	}
}

// jika kondisi pertama sudah terpenuhi maka hanya kondisi pertama yang dieksekusi
package main
import ("fmt")

func main() {
  x := 30

  if x >= 10 {
    fmt.Println("x is larger than or equal to 10.")
  } else if x > 20 {
    fmt.Println("x is larger than 20.")
  } else {
    fmt.Println("x is less than 10.")
  }
}

// Go Nested if Statement (if didalam if)
package main
import ("fmt")

func main()  {
	num:= 20

	if num > 10 {
		fmt.Println("Num is more than 10")
		if num > 15 {
			fmt.Println("Num is also more than 10")
		}
	} else{
		fmt.Println("Num is less than 10")
	}
}