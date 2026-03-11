package restapi

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Server!")
	})

	const port string = ":8080" // 127.0.0.1 = localhost (:8080 doesn't work for 127.0.0.1:8080 with WSL on Windows)

	fmt.Println("Server Listening on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("error starting server", err)
	}
}
