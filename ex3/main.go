package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleWelcome)
	http.HandleFunc("/welcome", handleWelcome)
	http.ListenAndServe(":8080", nil)

}
func handleWelcome(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	name = fmt.Sprintf("welocome %s \n", name)
	age = fmt.Sprintf("Age is %s \n", age)
	fmt.Fprint(w, name, age)

}
