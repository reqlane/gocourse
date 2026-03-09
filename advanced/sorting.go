package advanced

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

type person struct {
	name string
	age  int
}

type by func(p1, p2 *person) bool

type personSorter struct {
	people []person
	by     by
}

func (s *personSorter) Len() int {
	return len(s.people)
}
func (s *personSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

func (by by) sort(people []person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

func main() {

	// === sort.Sort() - not guaranteed stable
	people := []person{
		{"Alice", 30},
		{"Bob", 25},
		{"Ana", 35},
	}

	fmt.Println("Unsorted by age:", people)
	ageAsc := func(p1, p2 *person) bool {
		return p1.age < p2.age
	}
	ageDesc := func(p1, p2 *person) bool {
		return p1.age > p2.age
	}
	name := func(p1, p2 *person) bool {
		return p1.name < p2.name
	}
	lenName := func(p1, p2 *person) bool {
		return len(p1.name) < len(p2.name)
	}

	by(ageAsc).sort(people)
	fmt.Println("Sorted by age ascending:", people)
	by(ageDesc).sort(people)
	fmt.Println("Sorted by age descending:", people)
	by(name).sort(people)
	fmt.Println("Sorted by name:", people)
	by(lenName).sort(people)
	fmt.Println("Sorted by length of name:", people)

	// === sort.Slice() - not guaranteed stable
	stringSlice := []string{"guava", "apple", "cherry", "grapes", "banana"}
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})
	fmt.Println("Sorted by last character:", stringSlice)

	// === slices.Sort() - not guaranteed stable
	fmt.Println("slices package")

	stringSlice = []string{"guava", "apple", "Cherry", "grapes", "banana"}
	slices.Sort(stringSlice)
	fmt.Println("Sorted by slices.Sort():", stringSlice)

	// === slices.SortFunc() - not guaranteed stable
	people = []person{
		{"Alice", 30},
		{"Bob", 25},
		{"Ana", 35},
	}
	slices.SortFunc(people,
		func(p1, p2 person) int {
			return cmp.Compare(p1.name, p2.name)
		})
	fmt.Println("Sorted by name slices.SortFunc():", people)

	// === slices.SortStableFunc() - guaranteed stable
	slices.SortStableFunc(people,
		func(p1, p2 person) int {
			return cmp.Compare(p1.age, p2.age)
		})
	fmt.Println("Sorted by name slices.SortFunc():", people)
}
