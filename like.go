package main

type Person struct {
	Name  string
	Likes []string
}

var people []*Person

func initPeople(people []*Person) []*Person {
	names := []string{"a", "b", "c"}
	var p *Person
	for i := 0; i < 10; i++ {
		p = new(Person)
		p.Name = names[len(names)%3]
	}
	return people
}
