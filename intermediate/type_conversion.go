package intermediate

import "fmt"

func main() {

	var a int = 32
	b := int32(a)
	c := float64(b)
	// d := bool("correct")

	e := 3.14
	f := int(e)
	fmt.Println(f, c)

	// Type(value)
	g := "Hello @ привет"
	var h []byte
	h = []byte(g)
	fmt.Println(h)
	i := []byte{255, 72, 'a'}
	j := string(i)
	fmt.Println(j)
}
