// jika nilai variabel tidak berubah bisa mengguankan constans
const CONSTNAME type = value

// Declaring a Constant
package main
import ("fmt")

// nama variabel const biasanya menggunakan huruf kapital
const PI = 3.14 // bisa dideklarasikan di dalam atau di luar fungsi

func main(){
	// const PI = 3.14 // bisa dideklarasikan di dalam atau di luar fungsi
	fmt.Println(PI) // output: 3.14
}

// Typed Constants
// menyertakan tipe data
package main
import ("fmt")

const A int = 1

func main(){
	fmt.Println(A) // output: 1
}

// Untyped Constants
// tidak menyertakan tipe data
package main
import ("fmt")

const A = 1

func main(){
	fmt.Println(A) // output: 1
}

// Constants: Unchangeable and Read-only
package main
import ("fmt")

const A = 1
A = 2

func main(){
	fmt.Println(A) // cannot assign to A
}

// Multiple Constants
package main
import ("fmt")

const {
	A int = 1
	B = 3.14
	C = "Hi"
}

func main(){
	fmt.Println(A) // output: 1
	fmt.Println(B) // output: 3.14
	fmt.Println(C) // coutput: Hi
}