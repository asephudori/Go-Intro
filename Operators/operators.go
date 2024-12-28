// operators
package main
import("fmt")

func main() {
  var (
    sum1 = 100 + 50 // 150 (100 + 50)
    sum2 = sum1 + 250 // 400 (150 + 250)
    sum3 = sum2 + sum2 // 800 (400 + 400)
  )
  fmt.Println(sum3)
}

// arithmetic
package main
import ("fmt")

func main() {
  var (
	tambah = 1 + 1
	kurang = 2 - 1
	kali = 1 * 2
	bagi = 2 / 2
	modulus = 10 % 2
  )

  increment := 5
  decrement := 5
  increment++
  decrement--

  fmt.Println(tambah)
  fmt.Println(kurang)
  fmt.Println(kali)
  fmt.Println(bagi)
  fmt.Println(modulus)
  fmt.Println(increment)
  fmt.Println(decrement)
  /*output
	2
	1
	2
	1
	0
	6
	4
	*/
}

// Assignment Operators
package main
import ("fmt")

func main() {
  var x = 5
  x +=3 // x = x + 3
  x -=3 // x = x - 3
  x *= 3 // x = x * 3
  x /= 3 // x = x / 3
  x %= 3 // x = x % 3
  x &= 3 // x = x & 3
  x |= 3 // x = x | 3
  x ^= 3 // x = x ^ 3
  x >>= 3 // x = x >> 3
  x <<= 3 // x = x << 3

  fmt.Println(x)
}

// Comparison Operators
package main
import ("fmt")

func main() {
  var x = 5
  var y = 3

  fmt.Println(x>y)
  fmt.Println(x<y) 
  fmt.Println(x>=y) 
  fmt.Println(x<=y) 
  fmt.Println(x==y) 
  fmt.Println(x!=y)
}

// Logical Operators
package main
import ("fmt")

func main() {
  var x = 5

  fmt.Println(x < 5 &&  x < 10)
  fmt.Println(x < 5 || x < 4)
  fmt.Println(!(x < 5 && x < 10))
}

