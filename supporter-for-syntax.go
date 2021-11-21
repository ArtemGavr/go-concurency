package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	birthYear int
	name      string
}

func dummySwap(person *Person) (a string, b string) {

	convertedInt := strconv.Itoa(person.birthYear)
	return person.name, convertedInt
}

func mainSupporter() {
	defer fmt.Println("Goodbye")

	fmt.Println("Hello, 世界")

	phrases := [2]string{" was ", " lived "}

	var myName = "Artem"
	currentYear := 2002
	const yearBorn = 2002
	for i := 0; i < 10; i++ {
		name, year := dummySwap(&Person{name: myName, birthYear: yearBorn})

		phraseToUse := phrases[currentYear % 2]

		age := strconv.Itoa(currentYear - yearBorn)
		fmt.Println(name + " " + phraseToUse + age + " in " + year)

			currentYear++
	}


	//type BadOrder struct
	//{
	//	orderNum int
	//	customer string
	//}
	//
	//order1 := BadOrder{1, "Joe"}
	//order2 := BadOrder{2, "Joe"}
	//


	type Costumer struct
	{
		name string
	}
	type Order struct
	{
		  orderNum int
		  customer Costumer
	}

	costumer := Costumer{name:"Joe"}
	order1 := Order{1, costumer}
	order2 := Order{2, costumer}



	if order1 == order2 {}


}
