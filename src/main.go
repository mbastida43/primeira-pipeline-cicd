package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
<<<<<<< HEAD
	fmt.Fprintf(w, "Aplicacao exemplo - v2")
=======
	fmt.Fprintf(w, "Aplicacao exemplo")
>>>>>>> 8f17bb3 (commit 1st - pipeline CD)
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}
