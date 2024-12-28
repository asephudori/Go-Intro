// Slices, mirip seperti array namun lebih powerful dan flexible
// Create a Slice With []datatype{values}
// slice_name := []datatype{values}
package main

import (
	"fmt"
)

func main (){
	mySlice1 := []int{}
	fmt.Println(len(mySlice1)) // length: menampilkan jumlah karakter dalam slices
	fmt.Println(cap(mySlice1)) // capacity: menampilkan jumlah kapasitas element yang bisa bertambah/berkurang dalam slices
	fmt.Println(mySlice1)
	
	mySlice2 := []string{"Go","Slices","are", "Powerful"}
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))
	fmt.Println(mySlice2)
	// output
	// 0
	// 0
	// []
	// 4
	// 4
	// [Go Slices are Powerful]
}

// Create a Slice From an Array
/* var myarray = [length]datatype{values} // An array
myslice := myarray[start:end] // A slice made from the array */
package main
import ("fmt")

func main (){
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	mySlice := arr1[2:4]

	fmt.Printf("mySlice = %v\n", mySlice)
	fmt.Printf("length = %d\n", len(mySlice))
	fmt.Printf("capacity = %d\n", cap(mySlice))
	// output
	// mySlices = [12 13]
	// length = 2
	// capacity = 4
}

// Create a Slice With The make() Function
// slice_name := make([]type, length, capacity)
package main
import ("fmt")

func main (){
	mySlice1 := make([]int, 5, 10)

	fmt.Printf("My Slice 1 = %v\n", mySlice1)
	fmt.Printf("length = %d\n", len(mySlice1))
	fmt.Printf("capacity = %d\n", cap(mySlice1)) 
	
	mySlice2 := make([]int, 5) // jika tidak ada parameter untuk capacity maka akan mengikuti lenght
	
	
	fmt.Printf("My Slice 2 = %v\n", mySlice2)
	fmt.Printf("length = %d\n", len(mySlice2))
	fmt.Printf("capacity = %d\n", cap(mySlice2)) 
	// output
	// My Slice 1 = [0 0 0 0 0]
	// length = 5
	// capacity = 10
	// My Slice 2 = [0 0 0 0 0]
	// length = 5
	// capacity = 5
}

// Access Elements of a Slice
package main
import ("fmt")

func main()  {
	prices := []int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])
	/*output
	10
	30
	*/
}

// Change Elements of a Slice
package main
import ("fmt")

func main()  {
	prices := []int{10, 20, 30}
	prices[2] = 50

	fmt.Println(prices[0])
	fmt.Println(prices[2])
	/*output
	10
	50
	*/
}

// Append Elements To a Slice
// slice_name = append(slice_name, element1, element2, ...)
package main
import ("fmt")

func main()  {
	mySlice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("My Slice 1 = %v\n", mySlice1)
	fmt.Printf("Length = %d\n", len(mySlice1))
	fmt.Printf("Capacity = %v\n", cap(mySlice1))
	
	mySlice1 = append(mySlice1, 20, 21)
	fmt.Printf("My Slice 1 = %v\n", mySlice1)
	fmt.Printf("Length = %d\n", len(mySlice1))
	fmt.Printf("Capacity = %v\n", cap(mySlice1))
	/*output
	My Slice 1 = [1 2 3 4 5 6]
	Length = 6
	Capacity = 6
	My Slice 1 = [1 2 3 4 5 6 20 21]
	Length = 8
	Capacity = 12
	*/
}

// Append One Slice To Another Slice
// slice3 = append(slice1, slice2...)
// Note: The '...' after slice2 is necessary when appending the elements of one slice to another.

package main
import ("fmt")

func main()  {
	mySlice1 := []int{1, 2, 3}
	mySlice2 := []int{4, 5, 6}
	mySlice3 := append(mySlice1, mySlice2...)

	fmt.Printf("My Slice 3 = %v\n", mySlice3)
	fmt.Printf("Length = %d\n", len(mySlice3))
	fmt.Printf("Capacity = %d\n", cap(mySlice3))

	/* output
	My Slice 3 = [1 2 3 4 5 6]
	Length = 6
	Capacity = 6
	*/
}

// Change The Length of a Slice
package main
import ("fmt")

func main()  {
	arr1 := [6]int{9, 10, 11, 12, 13, 14}
	
	mySlice1 := arr1[1:5] // slice array
	
	fmt.Printf("My Slice 1 = %v\n", mySlice1)
	fmt.Printf("Length = %d\n", len(mySlice1))
	fmt.Printf("Capacity = %d\n", cap(mySlice1))
	
	mySlice1 = arr1[1:3] // mengubah length array dengan re-slicing
	
	fmt.Printf("My Slice 1 = %v\n", mySlice1)
	fmt.Printf("Length = %d\n", len(mySlice1))
	fmt.Printf("Capacity = %d\n", cap(mySlice1))
	
	mySlice1 = append(mySlice1, 20, 21, 22, 23) // mengubah length array dengan appending

	fmt.Printf("My Slice 1 = %v\n", mySlice1)
	fmt.Printf("Length = %d\n", len(mySlice1))
	fmt.Printf("Capacity = %d\n", cap(mySlice1))

	/* output
	My Slice 1 = [10 11 12 13]
	Length = 4
	Capacity = 5
	My Slice 1 = [10 11]
	Length = 2
	Capacity = 5
	My Slice 1 = [10 11 20 21 22 23]
	Length = 6
	Capacity = 10
	*/
}

// Memory Efficiency
/* When using slices, Go loads all the underlying elements into the memory.
 The copy() function creates a new underlying array with only the required elements for the slice.
 This will reduce the memory used for the program.
 copy(dest, src) */
package main
import ("fmt")

func main()  {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// original slices
	fmt.Printf("Numbers = %v\n", numbers)
	fmt.Printf("Length = %d\n", len(numbers))
	fmt.Printf("Capacity = %d\n", cap(numbers))
	
	// create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)
	
	fmt.Printf("Numbers Copy = %v\n", numbersCopy)
	fmt.Printf("Length = %d\n", len(numbersCopy))
	fmt.Printf("Capacity = %d\n", cap(numbersCopy))
	/*output
	Numbers = [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
	Length = 15
	Capacity = 15
	Numbers Copy = [1 2 3 4 5]
	Length = 5
	Capacity = 5
	*/
}