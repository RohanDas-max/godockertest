package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("my first docker go app")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "my first docker app")
	})

	log.Fatal(http.ListenAndServe(":4000", nil))
}
