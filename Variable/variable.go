// Declaring (Creating) Variables
// var
// bisa di deklarasikan di dalam atau di luar fungsi
var variableName type = value

// :=
// harus dalam satu baris
variableName := value

// Variable Declaration With Initial Value
package main
import ("fmt")

func main() {
	var student1 string = "Jhon" // tipe data string
	var student2 = "Jane" // tipe data ditentukan oleh compiler
	x := 2 // tipe data ditentukan oleh compiler

	fmt.Println(student1) // output: Jhon
	fmt.Println(student2) // output: Jane
	fmt.Println(x) // output: 2
}

//  Variable Declaration Without Initial Value
package main
import ("fmt")

func main(){
	var a string
	var b int
	var c bool

	fmt.Println(a) // output: ""
	fmt.Println(b) // output: 0
	fmt.Println(c) // output: false
}

// Value Assignment After Declaration
package main
import ("fmt")

func main(){
	var student1 string
	student1 = "Jhon"
	
	fmt.Println(student1) // output: Jhon
}

// declaring variables outside of a function
package main
import ("fmt")

var a int
var b int = 2
var c = 3

func main(){
	a = 1
	fmt.Println(a) // output: 1
	fmt.Println(b) // output: 2
	fmt.Println(c) // output: 3
}

// Multiple Variable Declaration
package main
import ("fmt")

func main(){
	// jika tipe data ditentukan langsung maka hanya satu tipe data yang bisa digunakan dalan satu baris
	var a, b, c, d int = 1, 3, 5, 7

	// jika tipe data tidak ditentukan maka dalam satu baris bisa menggunakan beberapa tipe data
	var a, b = 6, "Hello"
	c, d := 7, "World"

	fmt.Println(a) // output: 6
	fmt.Println(b) // output: Hello
	fmt.Println(c) // output: 7
	fmt.Println(d) // output: World
}

// Variable Declaration in a Block
package main
import ("fmt")

func main(){
	// deklarasi variabel dalam satu blok
	var (
		a int
		b int = 1
		c string = "hello"
	)

	fmt.Println(a) // output: 0
	fmt.Println(b) // output: 1
	fmt.Println(c) // output: hello
}