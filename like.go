package main

import "fmt"

type Person struct {
	Name  string
	Likes []string
}

var people []*Person

func initPeople(people []Person) []Person {
	names := []string{"a", "b", "c", "d", "e", "f", "g"}
	for i := 0; i < len(people); i++ {
		people[i].Name = names[i%7]
	}
	people[0].Likes = []string{"apple"}
	people[1].Likes = []string{"apple", "banana"}
	people[2].Likes = []string{"banana", "mango", "peach"}
	people[3].Likes = []string{"peach", "apple", "grape"}
	return people
}

func main() {
	people := []Person{{}, {}, {}, {}, {}, {}, {}, {}}
	people = initPeople(people)
	for _, p := range people {
		fmt.Println(p.Name)
	}

	likes := make(map[string][]Person)
	for _, p := range people {
		for _, l := range p.Likes {
			likes[l] = append(likes[l], p)
		}
	}
	for f, p := range likes {
		fmt.Print("\n", f, ":")
		for _, s := range p {
			fmt.Print(s.Name)
		}
	}
}
