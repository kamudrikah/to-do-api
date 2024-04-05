package main

import (
	"fmt"
	"kamudrikah/to-do-api/controllers"

	"github.com/gorilla/mux"
)

func addApproutes(route *mux.Router) {
	route.HandleFunc("/api/task", controllers.GetTasks).Methods("GET")
	route.HandleFunc("/api/task", controllers.InsertTask).Methods("POST")
	route.HandleFunc("/api/task", controllers.UpdateTask).Methods("PUT")
	route.HandleFunc("/api/task/{id}", controllers.DeleteTask).Methods("DELETE")

	fmt.Println("Routes are Loded.")
}
