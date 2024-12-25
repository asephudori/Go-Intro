package main
import ("fmt")

func main()
// akan error karena semicolon tidak boleh di awal baris
{ 
  fmt.Println("Hello World!")
}

// kode bisa dipersingkat dengan cara seperti ini, namun tidak disarankan karena akan sulit dibaca
package main; import ("fmt"); func main() { fmt.Println("Hello World!");}