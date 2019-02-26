package main

import "fmt"

type customer struct {
	id   int
	name string
}

func main() {
	var name = "a hello world"
	//prints ascii of the char
	fmt.Println(name[0])

	//different type of declarations
	x := 10
	fmt.Println(x)

	//the value will not be modified
	//by default the funcs use pass by value
	printStr(name)
	fmt.Println("String value after modification ", name)

	//the value is modified in this call
	printStrPointer(&name)

	fmt.Println("String value after modification using pointer ", name)

	createCustomer()
}

func printStrPointer(strPtr *string) {
	fmt.Println("Printing string pointer value --> ", *strPtr)
	fmt.Println("Printing string pointer address --> ", strPtr)

	fmt.Println("modifying string pointer value ")
	*strPtr = " Hello World"
}

func printStr(str string) {
	fmt.Println("Printing string value --> ", str)

	fmt.Println("modifying string value ")
	str = " Hello World"
}

func createCustomer() {
	//new returns *T, a pointer to the type
	cust := new(customer)
	cust.name = "God"
	cust.id = 1

	printCustomer(cust)
	modifyCustomer(cust)
	printCustomer(cust)

}

func printCustomer(cust *customer) {
	fmt.Println("Customer ", cust.id, " ", cust.name)
}

func modifyCustomer(cust *customer) {
	fmt.Println("Modifying customer id to 2")
	cust.id = 2
}
