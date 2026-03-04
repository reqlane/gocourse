package intermediate

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	user := os.Getenv("USERNAME")
	home := os.Getenv("HOME")

	fmt.Println("USERNAME environment variable:", user)
	fmt.Println("HOME environment variable:", home)

	// Set environment variable
	err := os.Setenv("FRUIT", "APPLE")
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
		return
	}
	fmt.Println("Set environment variable successfully on key FRUIT")

	fmt.Println("FRUIT:", os.Getenv("FRUIT"))

	for _, e := range os.Environ() {
		kvpair := strings.SplitN(e, "=", 2)
		fmt.Println(kvpair[0])
	}

	// Unset environment variable
	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error unsetting environment variable:", err)
		return
	}
	fmt.Println("Unset environment variable done on key FRUIT")
	fmt.Println("FRUIT:", os.Getenv("FRUIT"))

	fmt.Println("------------------------------------------")
	str := "a=b=c=d=e"
	fmt.Println(strings.SplitN(str, "=", -1))
	fmt.Println(strings.SplitN(str, "=", 0))
	fmt.Println(strings.SplitN(str, "=", 1))
	fmt.Println(strings.SplitN(str, "=", 2))
	fmt.Println(strings.SplitN(str, "=", 3))
	fmt.Println(strings.SplitN(str, "=", 4))
	fmt.Println(strings.SplitN(str, "=", 5))
	fmt.Println(strings.SplitN(str, "=", 6))
}
