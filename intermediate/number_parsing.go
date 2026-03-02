package intermediate

import (
	"fmt"
	"strconv"
)

func main() {

	// Atoi
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}
	fmt.Println("Parsed Integer:", num)
	fmt.Println("Parsed Integer:", num+1)

	// ParseInt
	numistr, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
		return
	}
	fmt.Println("Parsed Integer:", numistr)

	// ParseFloat
	floatStr := "3.14"
	floatVal, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Error parsing value:", err)
		return
	}
	fmt.Printf("Parsed float: %.2f\n", floatVal)

	// Binary
	binaryStr := "1010" // 0 + 2 + 0 + 8 = 10
	decimal, err := strconv.ParseInt(binaryStr, 2, 32)
	if err != nil {
		fmt.Println("Error parsing binary value:", err)
		return
	}
	fmt.Println("Parsed binary to decimal:", decimal)

	// Hex
	hexStr := "Ff" // 15 + 240 = 255
	hex, err := strconv.ParseInt(hexStr, 16, 32)
	if err != nil {
		fmt.Println("Error parsing hex value:", err)
		return
	}
	fmt.Println("Parsed hex to decimal:", hex)

	// Invalid number
	invalidNum := "456abc"
	invalidParsed, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Println("Error parsing number:", err)
		return
	}
	fmt.Println("Parsed invalid number:", invalidParsed)
}
