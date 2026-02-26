package basics

import "fmt"

func main() {

	// // Switch statement in Go (switch case default) (fallthrough)
	// switch expression {
	// case value1:
	// 	// code 1
	// fallthrough
	// case value2:
	// 	// code 2
	// case value3:
	// 	// code 3
	// default:
	// 	// default code if none of the cases are true
	// }

	// // Switch statement in other languages (switch case break default break)
	// switch expression {
	// case value1:
	// 	// code 1
	// break;
	// case value2:
	// 	// code 2
	// break;
	// case value3:
	// 	// code 3
	// break;
	// default:
	// 	// default code if none of the cases are true
	// break;
	// }

	fruit := "pineapple"
	switch fruit {
	case "apple":
		fmt.Println("It's an apple")
	case "banana":
		fmt.Println("It's banana")
	default:
		fmt.Println("Unknown fruit")
	}

	// Multiple conditions
	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday")
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend")
	default:
		fmt.Println("Invalid day")
	}

	// Expressions
	number := 15
	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or more")
	}

	// fallthrough
	num := 2
	switch {
	case num > 1:
		fmt.Println("Greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Not two")
	}

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's integer")
	case int32:
		fmt.Println("It's integer32")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's string")
	default:
		fmt.Println("Unknown type")
	}
}
