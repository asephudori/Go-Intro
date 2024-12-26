// The Print() Function
// menampilkan hasil sesuai format default
package main
import ("fmt")

func main(){
	var i,j string = "Hello","World"

	fmt.Print(i)
	fmt.Print(j)
	// output: HelloWorld
	// jika ingin menampilkan hasil di new line
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")
	/*	output: Hello
				World */
			}
			
// menggunakan satu argumen print untuk beberapa variabel
package main
import ("fmt")
			
func main(){
	var i,j string = "Hello","World" // spasi bisa ditambahkan dalam varibel
	var i,j = 10,20 // spasi akan ditambahkan otomatis selain tipe data string
				
	fmt.Print(i, "\n" ,j)
	fmt.Print(i, " " ,j) // spasi ditambahkan dalam argumen print
	/*	output: Hello
				World */
}

// The Println() Function
// akan menambahkan spasi diantara beberapa argumen dan menambahkan new line diakhir
package main
import ("fmt")

func main(){
	var i,j string = "Hello","World"

	fmt.Println(i,j)
	// output: Hello World
}

// The Printf() Function
// memberikan kita pilihan untuk menampilkan output sesuai format yang diberikan
// Format Specifier
package main
import "fmt"

func main() {
    angka := 123
    pecahan := 3.14159
    teks := "Hello"

    fmt.Printf("Angka: %d\n", angka)        // Output: Angka: 123
    fmt.Printf("Pecahan: %.2f\n", pecahan) // Output: Pecahan: 3.14
    fmt.Printf("Teks: %s\n", teks)          // Output: Teks: Hello
    fmt.Printf("Nilai default: %v\n", angka) // Output: Nilai default: 123
    fmt.Printf("Tipe data: %T\n", angka)   // Output: Tipe data: int
    fmt.Printf("Biner: %b\n", angka)        // Output: Biner: 1111011
    fmt.Printf("Oktal: %o\n", angka)        // Output: Oktal: 173
    fmt.Printf("Heksadesimal: %x\n", angka) // Output: Heksadesimal: 7b
    fmt.Printf("Persen: %%\n")            // Output: Persen: %
}

// Flag dan Lebar Field
package main
import "fmt"

func main() {
    angka := 42

    fmt.Printf("%10d\n", angka)   // Output:         42 (10 spasi di depan)
    fmt.Printf("%-10d\n", angka)  // Output: 42         (10 spasi di belakang)
    fmt.Printf("%010d\n", angka)  // Output: 0000000042 (diisi angka 0)
    fmt.Printf("%.3f\n", 3.14) // Output: 3.140
}

// formatting verbs
// https://www.w3schools.com/go/go_formatting_verbs.php,

// Unsigned Integers
// hanya menampilkan nilai non-negatif menggunakan keyword uint
package main
import ("fmt")

func main() {
  var x uint = 500
  var y uint = 4500
 
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}