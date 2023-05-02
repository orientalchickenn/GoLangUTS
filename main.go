package main

import (
	"Person"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 NOT FOUND", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Println(w, "ParseForm() err : %v", err)
			return
		}

		fmt.Fprintf(w, "Post from website r.postform = %v\n", r.PostForm)
		intVar, err := strconv.ParseInt((r.FormValue("age")), 0, 8)
		dataUser := Person.Orang{
			Name: r.FormValue("name"),
			Age:  intVar,
		}
		Use(err)

		fmt.Fprintf(w, "Your Name is %s\n", dataUser.Name)
		fmt.Fprintf(w, "You are %d years old\n", dataUser.Age)
		fmt.Fprintf(w, "You were born in %d\n", 2023-dataUser.Age)
	default:
		fmt.Fprintf(w, "Only GET and POST")
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Printf("Starting server!")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
