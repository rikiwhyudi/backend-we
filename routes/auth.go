package routes

import (
	"backend-WE/handlers"
	"backend-WE/pkg/mysql"
	"backend-WE/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
}
