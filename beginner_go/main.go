package main

import "fmt"

func main() {
	// Variables
	var name string = "Vinayak Vispute"
	fmt.Printf("This is my name %s\n", name)

	age := "21"
	fmt.Printf("%s's age is %d\n", name, age)

	var city string
	city = "Vadodara"
	fmt.Printf("%s lives in %s\n", name, city)

	var country, continent string = "India", "Asia"
	fmt.Printf("%s lives %s country which is located in %s continent\n", name, country, continent)

	var (
		isEmployed bool   = true
		salary     int    = 10000
		position   string = "developer"
	)
	fmt.Printf("isemployed : %t this is my salary: %d and this is my position %s", isEmployed, salary, position)

	var defaultInt int
	var defaultString string
	var defaultFloat float64
	var defaultBool bool

	fmt.Printf("defaultInt %d \ndefaultString '%s' \ndefaultFloat %f \ndefaultBool %t\n", defaultInt, defaultString, defaultFloat, defaultBool)

	const pi = 3.14

	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)

	fmt.Printf("Monday : %d, Tuesday : %d, Wednesdat : %d\n", Monday, Tuesday, Wednesday)

	const typedAge int = 20
	const unTypedAge = 20

	fmt.Println(typedAge == unTypedAge)

	const (
		JAN = iota + 1 // 1
		FEB            // 2
		MAR            // 3
		APR            // 4
	)

	fmt.Printf("Jan : %d, Feb : %d, March : %d, April : %d\n", JAN, FEB, MAR, APR)

	soln := add(1, 2)

	fmt.Printf("THE SUM IS : %d\n", soln)

	sum, product := calculateSumAndProduct(5, 2)

	fmt.Printf("THE SUM IS : %d & THIS IS PRODUCT : %d", sum, product)

}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func calculateSumAndProduct(num1, num2 int) (int, int) {
	return num1 + num2, num1 * num2
}
