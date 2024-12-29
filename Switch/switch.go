// Go switch Statement
// Single-Case switch
package main
import ("fmt")

func main()  {
	day := 4

	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturday")
	case 7:
		fmt.Println("sunday")
	}
}

/*
The default Keyword
The default keyword specifies some code to run if there is no case match:
*/
package main
import ("fmt")
	
func main()  {
	day := 8
	
	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturday")
	case 7:
		fmt.Println("sunday")
	default:
		fmt.Println("not a weekday")
	}
}

// Go Multi-case switch Statement
package main
import ("fmt")
	
func main()  {
	day := 5
	
	switch day {
	case 1,3,5:
		fmt.Println("odd weekday")
	case 2,4:
		fmt.Println("even weekday")
	case 6,7:
		fmt.Println("weekend")
	default:
		fmt.Println("invalid day of day number")
	}
}
