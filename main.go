package main

import (
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("John")

	http.Handle("/", templ.Handler(component))

	http.ListenAndServe(":3000", nil)
}
