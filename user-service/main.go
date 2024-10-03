package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = map[string]User{
	"1": {ID: "1", Name: "John Doe"},
	"2": {ID: "2", Name: "Jane Doe"},
}

// Получить пользователя по ID из URL
func getUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/")
	user, exists := users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Получить всех пользователей
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/user/", getUser)     // Для получения пользователя по ID
	http.HandleFunc("/users", getAllUsers) // Для получения всех пользователей
	http.ListenAndServe(":8084", nil)
}
