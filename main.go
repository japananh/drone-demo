package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Print(sum(2, 3))

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	fmt.Printf("Starting server at port 4444\n")
	if err := http.ListenAndServe("localhost:4444", nil); err != nil {
		log.Fatal(err)
	}
}

func sum(a, b int) int {
	return a + b
}
