package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "%s %s is %d years old!", user.FirstName, user.LastName, user.Age)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		user := User{
			FirstName: "John",
			LastName:  "Doe",
			Age:       30,
		}
		json.NewEncoder(w).Encode(user)
	})

	http.ListenAndServe(":8080", nil)
}

//$ curl -s -XPOST -d'{"firstname":"Elon","lastname":"Musk","age":48}' http://localhost:8080/decode
//Elon Musk is 48 years old!
//
//$ curl -s http://localhost:8080/encode
//{"firstname":"John","lastname":"Doe","age":25}
