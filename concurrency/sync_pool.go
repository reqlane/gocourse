package main

import (
	"fmt"
	"sync"
)

// Where to Use It
// Use sync.Pool when you are creating thousands of short-lived, identical objects per second.
// Best Example: A web server handling JSON requests.
// Instead of creating a new bytes.Buffer for every single request, you grab one from the pool, use it, and put it back.

// val := pool.Get()
// typeInstance, ok := val.(type)
// if !ok {
//     fmt.Println("Error: value was not of", type)
// }

// Where NOT to Use It
// Long-lived objects: If you have a database connection that
// stays open for an hour, don't pool it. Just keep the variable.
// Small, cheap objects: Don't pool a simple int or a small struct.
// The overhead of managing the pool is actually slower than just letting the GC handle a tiny 8-byte allocation.
// Stateful objects you forget to reset: If you put a "User" object in a pool with the name "Bob"
// and forget to clear it, the next person to Get() might accidentally see "Bob's" data.

type person struct {
	name string
	age  int
}

func main() {

	var pool = sync.Pool{
		New: func() any {
			fmt.Println("Creating a new person")
			return &person{}
		},
	}
	pool.Put(&person{name: "John", age: 81})

	// Get an object from the pool
	person1 := pool.Get().(*person)
	// person1.name = "John"
	// person1.age = 18
	// fmt.Println("Got person:", person1)

	fmt.Printf("Person 1. Name: %s, age: %d\n", person1.name, person1.age)

	pool.Put(person1)
	fmt.Println("Returned person to pool")

	person2 := pool.Get().(*person)
	fmt.Println("Got person 2:", person2)

	// person3 := pool.Get().(*person)
	person3 := pool.Get()
	if person3 != nil {
		fmt.Println("Got person 3:", person3)
		person3.(*person).name = "Jane"
		pool.Put(person3)
	} else {
		fmt.Println("Sync Pool is empty")
	}

	// Returning object to the pool
	pool.Put(person2)
	fmt.Println("Returned persons to pool")

	person4 := pool.Get().(*person)
	fmt.Println("Got person 4:", person4)

	person5 := pool.Get().(*person)
	fmt.Println("Got person 5:", person5)
}
