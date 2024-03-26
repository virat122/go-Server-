package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("starting server at port 8080 ")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parse form (), error :%v", err)
		return
	}

	fmt.Fprintf(w, "post request is sucessfull  \n")

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name is = %s\n", name)

	fmt.Fprintf(w, "address  is = %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found ", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "mothod not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello ")

}
