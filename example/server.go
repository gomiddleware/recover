package main

import (
	"math/rand"
	"net/http"

	"../"
)

func home(w http.ResponseWriter, r *http.Request) {
	// panic every 1 in 3 requests
	if rand.Intn(3) == 0 {
		panic("Waaaahhhh!")
	}
	w.Write([]byte(r.URL.Path))
}

func main() {
	home := http.HandlerFunc(home)

	http.Handle("/", recover.Handler(home))
	http.ListenAndServe(":8080", nil)
}
