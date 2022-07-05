package main

import (
	"encoding/json"
	"fmt"
	"io"
	"learngo/entity"
	"net/http"

	"github.com/go-chi/chi"
)

var person entity.Person

func main() {
	/*
			a, _ := add(1, 2)
			fmt.Println(a)
			workshop.Print("test")
			for test := 0; test < 10; test++ {
				fmt.Printf("I am number %v \n", test)
			}

			people := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			fmt.Print(people)

					for test := 0; test < 10; test++ {
						fmt.Printf("I am number %v \n", people[test])
					}

				for i, person := range people {
					fmt.Printf("I am person %v \n", person)
					fmt.Printf("I am number %v \n", people[i])
				}

		var name string
		fmt.Scanln(&name)
		var newIntern = intern{
			name:   name,
			gender: "Male",
			age:    30,
		fmt.Println(newIntern)}
	*/
	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		r.Get("/person", GetPerson)
		r.Post("/person", CreatePerson)

	})
	router.Route("/v2", func(r chi.Router) {
		r.Get("/person", GetPerson)
	})

	http.ListenAndServe(":8080", router)

}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	value := fmt.Sprintf("Person is %v", person)
	fmt.Fprint(w, value)

}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person = entity.Person{
		Name:   "John",
		ID:     "123",
		Gender: "Male",
	}
	body := r.Body
	decodedBody, _ := io.ReadAll(body)
	json.Unmarshal(decodedBody, &person)
	value := fmt.Sprintf("My name is %s", person)
	fmt.Fprintf(w, value)

}
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	name := r.URL.Query().Get("name")
	name = fmt.Sprintf("My name is %s", name)
	fmt.Fprint(w, name)

}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	value := fmt.Sprintf("Person is %v", person)
	fmt.Fprint(w, value)

}

/*
func add(a int, b int) (int, int) {
	return a + b, a - b
}

type intern struct {
	name   string
	gender string
	age    int
}
*/
