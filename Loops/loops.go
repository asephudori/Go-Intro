// Go For Loops
package main
import ("fmt")

func main()  {
	for i:=0; i < 100; i+=10{
		fmt.Println(i)
	}
}

// continue
// berfungsi utntuk melewati 1 atau lebih iterasi dalam loop
package main
import ("fmt")

func main()  {
	for i:=0; i < 5; i++{
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}

// The break Statement
// digunakan untuk break/terminate loop
package main
import ("fmt")

func main()  {
	for i:=0; i < 5; i++{
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}

// Nested Loops
// loop dalam loop
package main
import ("fmt")

func main()  {
	adj := [2]string{"big","tasty"}
	fruits := [3]string{"apple","orange","banana"}

	for i:=0; i < len(adj); i++ {
		for j:=0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}
}

// The Range Keyword
// menampilkan index dan nilai bersamaan
package main
import ("fmt")

func main()  {
	fruits := [3]string{"apple","orange","banana"}

	for idx, val := range fruits { // untuk menampilkan salah satu idx atau value ganti dengan (_) ex for _, val
		fmt.Printf("%v\t%v\n", idx, val)
	}
}
