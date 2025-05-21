package basics

import "fmt"

func main() {

	// Switch statement
	// switch expression {
	// case value1:
	// 	// Code to be executed if expression equals value1
	// case value2:
	// 	// Code to be executed if expression equals value2
	// case value3:
	// 	// Code to be executed if expression equals value3
	// default:
	// Code to be executed if expression does not match any value
	//}

	// fruit := "pineapple"

	// switch fruit {
	// case "apple":
	// 	fmt.Println("It's an apple.")
	// case "banana":
	// 	fmt.Println("It's a banana")
	// default:
	// 	fmt.Println("Unknown fruit.")
	// }

	// // Multiple Conditions
	// day := "Monday"

	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday")
	// case "Sunday":
	// 	fmt.Println("It's Sunday")
	// default:
	// 	fmt.Println("It's Saturday")
	// }

	// number := 15

	// switch {
	// case number < 10:
	// 	fmt.Println("Number is less than 10.")
	// case number >= 10 && number < 20:
	// 	fmt.Println("Number is between 10 and 20.")
	// default:
	// 	fmt.Println("Number is 20 or higher.")
	// }

	// num := 2

	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("Number is Two")
	// default:
	// 	fmt.Println("Not 2")
	// }

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case int32:
		fmt.Println("It's a 32bit integer")
	case float64:
		fmt.Println("It's a float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown Type")
	}
}
