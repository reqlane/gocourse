package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message1 := "Hello, \nGo!"
	message2 := "Hello, \tGo!"
	message3 := "Hello, Go!" // Go!lo,
	rawMessage := `Hello\nGo`

	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println(rawMessage)

	fmt.Println("Length of rawMessage variable is:", len(rawMessage))

	fmt.Println("The first character in message1 var is:", message1[0]) // ASCII

	greeting := "Hello "
	name := "Alice"
	fmt.Println(greeting + name)

	str1 := "Apple"  // A ASCII = 65
	str := "apple"   // a ASCII = 97
	str2 := "banana" // b ASCII = 98
	str3 := "app"    // a ASCII = 97
	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	for _, char := range message1 {
		// fmt.Printf("Character at i %d is %c\n", i, char)
		fmt.Printf("%v\n", char)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	jch := '日'

	fmt.Println(ch)
	fmt.Println(jch)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of cstr is %T\n", cstr)

	const NIHONGO = "日本語" // Japanese text
	fmt.Println(NIHONGO)

	jhello := "こんにちは" // Japanese "Hello"
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue)
	}

	r := '😊'
	fmt.Printf("%v\n", r)
	fmt.Printf("%c\n", r)
}
