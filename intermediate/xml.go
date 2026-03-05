package intermediate

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	Address Address  `xml:"address"`
	Email   string   `xml:"-"`
	// Email   string   `xml:"email"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state,attr"`
}

// <book isbn="qwerty123456" ...></book>
type Book struct {
	XMLName xml.Name `xml:"book"`
	ISBN    string   `xml:"isbn,attr"`
	Title   string   `xml:"title,attr"`
	Author  string   `xml:"author,attr"`
	Pseudo  Pseudo   `xml:"pseudo"`
}

type Pseudo struct {
	XMLName    xml.Name `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudoattr,attr"`
}

func main() {

	// Marshal
	// person := Person{Name: "Vladyslav", Age: 22, City: "Kyiv", Email: "vlad5000191@gmail.com"}
	person := Person{Name: "Vladyslav", Email: "vlad5000191@gmail.com", Address: Address{City: "Kyiv", State: "UA"}}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatalln("Error marshalling to xml:", err)
	}
	fmt.Println(string(xmlData))

	xmlData1, err := xml.MarshalIndent(person, "", "\t")
	if err != nil {
		log.Fatalln("Error marshalling to xml:", err)
	}
	fmt.Println(string(xmlData1))

	// Unmarshal
	// xmlRaw := `<person><name>Vladyslav</name><age>22</age></person>`
	xmlRaw := `<person><name>Vladyslav</name><age>22</age><address><city>Kyiv</city><state>UA</state></address></person>`

	var personXML Person

	err = xml.Unmarshal([]byte(xmlRaw), &personXML)
	if err != nil {
		log.Fatalln("Error unmarshalling XML:", err)
	}
	fmt.Println(personXML)
	fmt.Println("Local string:", personXML.XMLName.Local)
	fmt.Println("Namespace:", personXML.XMLName.Space)

	// XML attributes
	book := Book{
		ISBN:   "586-234-234-6543",
		Title:  "Go Bootcamp",
		Author: "Vladyslav",
		Pseudo: Pseudo{PseudoAttr: "PseudoAttr"},
	}

	xmlDataAttr, err := xml.MarshalIndent(book, "", "\t")
	if err != nil {
		log.Fatalln("Error marshalling data:", err)
	}
	fmt.Println(string(xmlDataAttr))
}
