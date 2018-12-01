package main

import (
	"fmt"
	"net/http"
)

// HELLOWORLD_START OMIT
func HelloWorld(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, World!")
}

// HELLOWORLD_END OMIT
