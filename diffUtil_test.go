package gotool

import (
	"fmt"
	"testing"
)

func Test_compareStructs(t *testing.T) {
	person := &Person{
		Name:             "Alice",
		Age:              25,
		Address:          &Address{City: "New York", State: "NY"},
		GoodNotesAbcHaha: "123",
	}
	employee := &Employee{
		Name:     "Alice",
		Age:      25,
		Salary:   5000,
		Address:  &Address{City: "San Francisco", State: "CA"},
		Position: "Manager",
	}
	fmt.Println("================================DiffModelDefault")
	from, to, err := DiffAndToJson(person, employee, DiffModelDefault, "address.state")
	if err != nil {
		return
	}
	fmt.Println("person  ", from)
	fmt.Println("employee", to)
	fmt.Println("================================DiffModelCareFromNotZeroValue")
	from, to, err = DiffAndToJson(person, employee, DiffModelCareFromNotZeroValue, "address.state")
	if err != nil {
		return
	}
	fmt.Println("person  ", from)
	fmt.Println("employee", to)
	fmt.Println("================================DiffModelCareToNotZeroValue")
	from, to, err = DiffAndToJson(person, employee, DiffModelCareToNotZeroValue, "address.state")
	if err != nil {
		return
	}
	fmt.Println("person  ", from)
	fmt.Println("employee", to)
}

type Address struct {
	City  string
	State string
}

type Person struct {
	Name             string
	Age              int
	Gender           string
	Address          *Address
	GoodNotesAbcHaha string
}

type Employee struct {
	Name     string
	Age      int
	Salary   int
	Address  *Address
	Position string
}
