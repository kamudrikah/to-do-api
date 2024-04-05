package main

import (
	"fmt"
	"kamudrikah/to-do-api/services"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8000"
	}

	fmt.Println("Server will start at http://localhost:" + appPort + "/")

	services.ConnectDatabse()

	route := mux.NewRouter()

	addApproutes(route)

	log.Fatal(http.ListenAndServe(":"+appPort, route))
}
