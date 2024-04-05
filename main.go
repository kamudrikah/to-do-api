package main

import (
	"fmt"
	"kamudrikah/to-do-api/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Server will start at http://localhost:8000/")

	services.ConnectDatabse()

	route := mux.NewRouter()

	addApproutes(route)

	log.Fatal(http.ListenAndServe(":8000", route))
}
