package intermediate

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	tempFileName := "temporaryFile"
	tempDirName := "GoCourseTempDir"
	tempFile, err := os.CreateTemp("", tempFileName)
	checkError(err)

	fmt.Println("Temporary file created:", tempFile.Name())

	defer func() {
		checkError(tempFile.Close())
		checkError(os.Remove(tempFile.Name()))
	}()

	tempDir, err := os.MkdirTemp("", tempDirName)
	checkError(err)

	defer func() {
		checkError(os.RemoveAll(tempDir))
	}()

	fmt.Println("Temporary directory created:", tempDir)

	i := 0

	defer func() {
		fmt.Println(i) // 1
	}()

	defer func(i int) {
		fmt.Println(i) // 0
	}(i)

	i++
}
