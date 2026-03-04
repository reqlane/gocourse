package intermediate

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed embed_directive.go
var content string

//go:embed basics
var basicsFolder embed.FS

func main() {

	fmt.Println("Embedded content:", content)
	content, err := basicsFolder.ReadFile("basics/import.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Embedded file content:", string(content))

	err = fs.WalkDir(basicsFolder, "basics", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error inside WalkDirFunc:", err)
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
