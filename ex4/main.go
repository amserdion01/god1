package main

import (
	"ex4/entity"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

var person = entity.Person{
	Name:   "John",
	ID:     123,
	Gender: "Male",
}

func main() {
	router := chi.NewRouter()
	router.Get("/person", GetPerson)
	http.ListenAndServe(":8080", router)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	value := fmt.Sprintf("Person is %v", person)
	fmt.Fprint(w, value)

}
