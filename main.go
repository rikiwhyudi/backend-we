package main

import (
	"fmt"
	"net/http"

	"backend-WE/database"
	"backend-WE/pkg/mysql"
	"backend-WE/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// env
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("server running localhost:5000")
	http.ListenAndServe(":5000", r)
}
