package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seconds := vars["seconds"]
	i, err := strconv.Atoi(seconds)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Fprintf(w, "Worked! for %d seconds :)", i)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{seconds}", handler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
