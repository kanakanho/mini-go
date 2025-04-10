package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kajiLabTeam/mr-platform-user-management-server/common"
	"github.com/kajiLabTeam/mr-platform-user-management-server/model"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user common.UserId
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Error decoding user ID:", err)
		return
	}

	// ユーザーが存在するか確認
	exist, err := model.ExistUser(user.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error checking user existence:", err)
		return
	}
	if exist {
		http.Error(w, "User already exists", http.StatusConflict)
		log.Println("User already exists:", user.UserId)
		return
	}

	// ユーザーを削除
	isDeleted, err := model.DeleteUser(user.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error deleting user:", err)
		return
	}
	if !isDeleted {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		log.Println("Failed to delete user:", user.UserId)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("User deleted successfully"))
}
