package intermediate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	// Length
	str := "Hello Go!"
	fmt.Println(len(str))

	// Concatenation
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	// Accessing and slicing
	fmt.Println(str[0])
	fmt.Println(str[1:7])

	// String Conversion
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(len(str3))

	// Strings splitting
	fruits := "apple, orange, banana"
	fruits1 := "apple-orange-banana"
	parts := strings.Split(fruits, ", ")
	parts1 := strings.Split(fruits1, "-")
	fmt.Println(parts)
	fmt.Println(parts1)

	// Strings joining
	countries := []string{"Germany", "France", "Italy"}
	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	// Contains
	fmt.Println(strings.Contains(str, "Go?"))

	// Replace / ReplaceAll
	replaced := strings.Replace(str, "Go", "Universe", 1)
	fmt.Println(replaced)
	replaced = strings.ReplaceAll(str, "Go", "Universe")
	fmt.Println(replaced)

	// Trim
	strwspace := " \tHello Everyone "
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace))

	// Lower/Upper case
	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))

	// Repeat
	fmt.Println(strings.Repeat("foo ", 3))

	// Count, Has Prefix/Suffix
	fmt.Println(strings.Count("Hello", "l"))
	fmt.Println(strings.HasPrefix("Hello", "he"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))
	fmt.Println(strings.HasSuffix("Hello", "la"))

	// Regexp
	str5 := "Hel1lo, 123 Go 11!"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches)

	// utf8
	str6 := "Hello, 世界"
	fmt.Println(utf8.RuneCountInString(str6))

	// STRING BUILDER
	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	// Convert builder to a string
	result = builder.String()
	fmt.Println(result)

	// Using WriteRune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you")

	result = builder.String()
	fmt.Println(result)

	// Reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh!")
	result = builder.String()
	fmt.Println(result)
}
