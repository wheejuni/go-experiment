package CustomStruct

import "fmt"

type Person struct {
	Name string
	Age int
	Profession int
}

func (p Person) callJobId() int {
	return p.Profession
}

func (p Person) String() string {
	return fmt.Sprintf("name is %s, age %d", p.Name, p.Age)
}

const(
	Programmer = iota
	Gambler
	BaseballPlayer
)