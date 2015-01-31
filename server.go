package main

import (
	"log"
	"net/http"
	"fmt"
)

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func (s Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s.Greeting)
	fmt.Fprint(w, s.Punct)
	fmt.Fprint(w, s.Who)
}

func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Humans!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
