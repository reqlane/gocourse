package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName    string  `json:"name"`
	Age          int     `json:"age,omitempty"`
	EmailAddress string  `json:"email,omitempty"`
	Address      Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}

func main() {

	person := Person{FirstName: "John"}

	// Marshal
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to json:", err)
		return
	}
	fmt.Println(string(jsonData))

	person1 := Person{FirstName: "Jane", Age: 30, EmailAddress: "jane@fakemail.com", Address: Address{City: "New York", State: "USA"}}

	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to json:", err)
		return
	}
	fmt.Println(string(jsonData1))

	// Unmarshal
	jsonData2 := `{"full_name": "Jenny Doe", "emp_id": "0009", "age": 30, "address": {"city": "San Jose", "state": "CA"}}`

	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonData2), &employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return
	}
	fmt.Println(employeeFromJson)
	fmt.Println("Jenny's age increased by 5 years:", employeeFromJson.Age+5)
	fmt.Println("Jenny's city:", employeeFromJson.Address.City)

	// Marshal list
	listOfCityState := []Address{
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
		{City: "Las Vegas", State: "NV"},
		{City: "Modesto", State: "CA"},
		{City: "Clearwater", State: "FL"},
	}

	fmt.Println(listOfCityState)
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		log.Fatalln("Error marshalling to json:", err)
	}
	fmt.Println("Json list:", string(jsonList))

	// Handling unknown json structures
	jsonData3 := `{"name": "John", "age": 30, "address": {"city": "New York", "state": "NY"}}`

	var data map[string]any

	err = json.Unmarshal([]byte(jsonData3), &data)
	if err != nil {
		log.Fatalln("Error unmarshalling json:", err)
	}
	fmt.Println("Decoded/Unmarshalled json:", data)
	fmt.Println("Decoded/Unmarshalled json:", data["address"])
	fmt.Println("Decoded/Unmarshalled json:", data["name"])
}
