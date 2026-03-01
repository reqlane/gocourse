package intermediate

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math Error: square root of negative number")
	}
	// compute the square root
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: empty data")
	}
	// process data
	return nil
}

func main() {

	// result, err := sqrt(16)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)

	// result, err = sqrt(-16)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)

	// data := []byte{}
	// // if err := process(data); err != nil {
	// err := process(data)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Data processed successfully")

	// error interface of builtin package
	// err := eprocess()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// println("")

	if err := readData(); err != nil {
		fmt.Println(err)
		fmt.Printf("The type is: %T\n", errors.Unwrap(err))
		return
	}
	fmt.Println("Data read successfully.")
}

type myError struct {
	message string
}

func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eprocess() error {
	return &myError{"Custom error message"}
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func readConfig() error {
	return &myError{"config error"}
}
