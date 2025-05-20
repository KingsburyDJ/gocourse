package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Eg. CalculateArea, UserInfo, NewHTTPRequest
	// Structs, interfaces, enums

	// snake_case (file names, some variables)
	// Eg. user_id, first_name, http_request

	// UPPERCASE (Constants!)

	// mixedCase (variables, identifiers from other languages)
	// Eg. javaScript, htmlDocument, isValid

	// lowercase (package names)

	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}
