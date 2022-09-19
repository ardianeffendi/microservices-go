package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s\n", d)
	})

	http.HandleFunc("/welcome", func(http.ResponseWriter, *http.Request) {
		log.Println("Welcome to \"welcome!\"")
	})

	http.ListenAndServe(":9090", nil)
}