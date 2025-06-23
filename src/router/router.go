package router

import (
	"log"
	"net/http"

	"github.com/kanakanho/mini-go/controller"
	_ "github.com/lib/pq"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		controller.CreateUser(w, r)
	case http.MethodGet:
		controller.GetUser(w, r)
	case http.MethodDelete:
		controller.DeleteUser(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Println("Method Not Allowed")
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func Init() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/hello", helloHandler)
	mux.HandleFunc("/api/user", userHandler)

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
