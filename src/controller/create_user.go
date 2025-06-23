package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kanakanho/mini-go/common"
	"github.com/kanakanho/mini-go/model"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	// ユーザーを作成
	isCreated, err := model.CreateUser(user.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error creating user:", err)
		return
	}
	if !isCreated {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		log.Println("Failed to create user:", user.UserId)
		return
	}

	userIdJson, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error marshalling user ID:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userIdJson)
}
