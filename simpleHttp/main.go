package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

type User struct {
	Id    int    `json:"Id"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := User{
		Id:    1,
		Name:  "test",
		Email: "test@gmail.com",
	}

	json.NewEncoder(w).Encode(user)
}
