package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kajiLabTeam/mr-platform-user-management-server/model"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	users, err := model.AllUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error fetching users:", err)
		return
	}
	if len(users.Users) == 0 {
		http.Error(w, "No users found", http.StatusNotFound)
		log.Println("No users found")
		return
	}
	usersJson, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error marshalling users:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usersJson)
	log.Println("Users fetched successfully:", users.Users)
}
