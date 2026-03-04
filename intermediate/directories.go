package intermediate

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		panic(err)
		// fmt.Println(err)
	}
}

func main() {

	// os.Mkdir
	checkError(os.Mkdir("subdir", 0755))
	defer func() {
		checkError(os.RemoveAll("subdir"))
	}()

	// os.WriteFile
	os.WriteFile("subdir/file", []byte(""), 0755)

	checkError(os.MkdirAll("subdir/parent/child", 0755))
	checkError(os.MkdirAll("subdir/parent/child1", 0755))
	checkError(os.MkdirAll("subdir/parent/child2", 0755))
	checkError(os.MkdirAll("subdir/parent/child3", 0755))
	os.WriteFile("subdir/parent/file", []byte(""), 0755)
	os.WriteFile("subdir/parent/child/file", []byte(""), 0755)

	result, err := os.ReadDir("subdir/parent")
	checkError(err)

	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	checkError(os.Chdir("subdir/parent/child"))
	dir, err := os.Getwd()
	checkError(err)
	fmt.Println(dir)

	result, err = os.ReadDir(".")
	checkError(err)

	fmt.Println("Reading subdir/parent/child...")
	for _, entry := range result {
		fmt.Println(entry)
	}

	checkError(os.Chdir("../../.."))
	dir, err = os.Getwd()
	checkError(err)
	fmt.Println(dir)

	// filepath.Walk and filepath.WalkDir
	pathfile := "subdir"
	fmt.Println("Walking directory...")
	err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		fmt.Println(path)
		return nil
	})
	checkError(err)

	checkError(os.Remove("subdir/file"))
	// checkError(os.RemoveAll("subdir"))
}
