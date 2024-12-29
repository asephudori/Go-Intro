// Go Functions
// create and call a fuction
package main
import ("fmt")

// create
func myMessage()  {
	fmt.Println("I just got executed")
}

func main()  {
	myMessage() // call
	myMessage() // bisa dipanggil lebih dari 1 kali
	myMessage()
}

// Go Function Parameters and Arguments
// When a parameter is passed to the function, it is called an argument.
package main
import ("fmt")

func familyName(fname string, lname string)  {
	fmt.Println("Hello", fname, "D", lname) // fname dan lname adalah parameter
}

func main(){
	familyName("Monkey","Luffy") // nama untuk mengisi fname dan lname adalah argumen
	familyName("Jeki","Shanks")
	familyName("Rora","Babymonster")
}

// Go Function Returns
// Return Values
package main
import ("fmt")

func myFunction(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(myFunction(1, 2))
}

// Named Return Values
package main
import ("fmt")

func myFunction(x int, y int) (result int) {
  result = x + y
  return // bisa tanpa nama variabel atau dengan nama variabel ex return result
}

func main() {
  fmt.Println(myFunction(1, 2))
}

// Store the Return Value in a Variable
package main
import ("fmt")

func myFunction(x int, y int) (result int) {
  result = x + y
  return
}

func main() {
  total := myFunction(1, 2)
  fmt.Println(total)
}

// Multiple Return Values
package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
  fmt.Println(myFunction(5, "Hello"))
  //dengan 2 variabel
  a, b := myFunction(5, "Hello")
  fmt.Println(a, b)
  // untuk menghilangkan salah satu nilai
  _, b := myFunction(5, "Hello")
  fmt.Println(b)
}

// Go Recursion Functions
package main
import ("fmt")

func factorial_recursion(x float64) (y float64) {
  if x > 0 {
     y = x * factorial_recursion(x-1)
  } else {
     y = 1
  }
  return
}

func main() {
  fmt.Println(factorial_recursion(4))
}