package intermediate

import "fmt"

func main() {

	var ptr *int
	var a int = 10
	ptr = &a // referencing

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(ptr)
	fmt.Println(*ptr) // dereferencing a pointer

	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int) {
	*ptr++
}
