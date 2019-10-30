package main

import "./CustomStruct"

func main() {
	println("go")

	sum, i := 0, 0

	for i < 10 {
		sum += i
		i++
	}

	println(sum)


	person := CustomStruct.Person{Profession:CustomStruct.Programmer, Name:"Zella", Age:47}
	println(person.String())


}