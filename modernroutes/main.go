package main

import (
	"fmt"
	"net/http"
)

// Go version 1.22+ required
func main() {

	mux := http.NewServeMux()

	// Method based routing
	mux.HandleFunc("POST /items/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Item created")
	})

	mux.HandleFunc("DELETE /items/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Item deleted")
	})

	// Wildcard in pattern - path parameter
	mux.HandleFunc("GET /teachers/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Teacher ID: %s", r.PathValue("id"))
	})

	// Wildcard with "..."
	// /files/part1/part2/part3/ ---> path=part1/part2/part3/
	mux.HandleFunc("/files/{path...}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Path: %s", r.PathValue("path"))
	})

	mux.HandleFunc("/path1/{param1}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Param: %s", r.PathValue("param1"))
	})

	// Neither is more specific error
	// mux.HandleFunc("/{param2}/path2", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Param: %s", r.PathValue("param2"))
	// })

	mux.HandleFunc("/path1/path2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Nothing to see here")
	})

	http.ListenAndServe(":8080", mux)
}
