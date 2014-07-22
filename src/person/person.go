package person

import "strconv"

type Person struct {
	Name string
	Age  int
}

func (p *Person) GetInfo() (info string) {
	info = "Name: " + p.Name + ", Age: " + strconv.Itoa(p.Age)
	return
}

func New(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}
