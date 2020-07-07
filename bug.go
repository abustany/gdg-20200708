package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Email string
	Name  string
}

func loadUsers() []User {
	// Let's assume this loads users from somewhere
	return []User{
		{Email: "alice@example.com", Name: "Alice"},
		{Email: "bob@example.com", Name: "Bob"},
		{Email: "mallory@example.com", Name: "Mallory"},
	}
}

func main() {
	usersByEmail := map[string]*User{}

	userList := loadUsers()

	for _, user := range userList {
		usersByEmail[user.Email] = &user
	}

	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		email := strings.TrimPrefix(r.URL.Path, "/user/")
		user, ok := usersByEmail[email]

		if !ok {
			http.Error(w, "could not find a user with this email", http.StatusNotFound)
			return
		}

		io.WriteString(w, "Hello "+user.Name+"!\n")
	})

	const listenAddress = "127.0.0.1:1412"

	log.Printf("Starting server on %s", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
