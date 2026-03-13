package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	// json.Marshal()
	user1 := user{Name: "Alice", Email: "alice@example.com"}
	fmt.Println(user1)
	jsonData1, err := json.Marshal(user1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonData1))

	// json.Unmarshal()
	var user2 user
	err = json.Unmarshal(jsonData1, &user2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("User created from json data:", user2)

	// json.NewDecoder(r io.Reader)
	jsonData2 := `{"name": "John", "email": "john@example.com"}`
	reader := strings.NewReader(jsonData2)
	decoder := json.NewDecoder(reader)

	var user3 user
	err = decoder.Decode(&user3)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user3)

	// json.NewEncoder(w io.Writer)
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)

	err = encoder.Encode(user1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Encoded json string:", buf.String())
}
