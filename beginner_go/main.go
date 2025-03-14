package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	person := Person{Name: "Vinayak Vispute", Age: 21}

	fmt.Printf("This is my person %v\n", person)
	fmt.Printf("This is my person %+v\n", person)

	employee := struct {
		Name string
		id   int
	}{
		Name: "Vadodara",
		id:   123423,
	}

	fmt.Printf("This is my employee : %+v\n", employee)

	type Address struct {
		Street string
		City   string
	}

	type Contact struct {
		Name    string
		Address Address
		Phone   string
	}

	contact := Contact{
		Name: "Vinayak",
		Address: Address{
			Street: "Tarsali Street",
			City:   "Vadodara",
		},
		Phone: "8401282182",
	}

	fmt.Printf("This is my Contact details : %+v\n", contact)

	fmt.Printf("Name before : %s\n", person.Name)
	modifyPersonName(&person)
	fmt.Printf("Name after : %s\n", person.Name)
	person.modifyPersonName2()
	fmt.Printf("Name after : %s\n", person.Name)

	x := 20
	ptr := &x
	fmt.Printf("Value of x: %d and address of x %p\n", x, ptr)
	*ptr = 30
	fmt.Printf("New Value of  x: %d and address of x %p\n", x, ptr)
}

func modifyPersonName(person *Person) {
	person.Name = "ModifiedName"
}

func (person *Person) modifyPersonName2() {
	person.Name = "SecondTimeNameModified"
}
