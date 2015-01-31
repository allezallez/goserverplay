package main

import (
	"log"
	"net/http"
	"fmt"
)

type String string

func (s String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s)
}

// this is a stupid struct
type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

//blah blah blah
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
