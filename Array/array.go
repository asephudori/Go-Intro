// Declare an Array
// With the var keyword
var array_name = [length]datatype{values} // here length is defined

or

var array_name = [...]datatype{values} // here length is inferred

// With the := sign
array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred

// defined
package main
import ("fmt")

func main ()  {
	var arr1 = [3]int{1,2,3}
	arr2 := [5]int{4,5,6,7,8}

	fmt.Println(arr1)
	fmt.Println(arr2)
	// output
	// [1 2 3]
	// [4 5 6 7 8]
}

// inferred 
package main
import ("fmt")

func main ()  {
	var arr1 = [...]int{1,2,3}
	arr2 := [...]int{4,5,6,7,8}

	fmt.Println(arr1)
	fmt.Println(arr2)
	// output
	// [1 2 3]
	// [4 5 6 7 8]
}

// Access Elements of an Array
package main
import ("fmt")

func main ()  {
	var arr1 = [...]int{1,2,3}
	arr2 := [...]int{4,5,6,7,8}

	fmt.Println(arr1[0])
	fmt.Println(arr2[4])
	// output
	// 1
	// 8
}
	
// Change Elements of an Array
package main
import ("fmt")
	
func main ()  {
	var arr1 = [...]int{1,2,3}
	arr2 := [...]int{4,5,6,7,8}

	arr1[2] = 9
	
	fmt.Println(arr1)
	fmt.Println(arr2[4])
	// output
	// [1,2,9]
	// 8
}

// Array Initialization
// jika kosong akan diisi otomatis dengan 0 atau ""
package main
import ("fmt")
	
func main ()  {
	var arr1 = [5]int{} // tidak di inisialisasi
	arr2 := [5]int{1,2} // sebagian
	arr3 := [5]int{1,2,3,4,5} // full
	
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	// output
	// [0,0,0,0,0]
	// [1,2,0,0,0]
	// [1,2,3,4,5]
}

// Initialize Only Specific Elements
package main
import ("fmt")
	
func main ()  {
	var arr1 = [5]int{1:10,2:40}

	fmt.Println(arr1)
	// output
	// [0,10,40,0,0]
}

// Find the Length of an Array
package main
import ("fmt")
	
func main ()  {
	var arr1 = [4]string{"Samsung","Pixel","Iphone","Nothing"}
	arr2 := [...]int{1,2,3,4,5,6}

	fmt.Println(arr1)
	fmt.Println(arr2)
	// output
	// 4
	// 6
}
