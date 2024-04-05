package main

import (
	"fmt"
	"kamudrikah/to-do-api/services"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	fmt.Println("Server will start at http://localhost:" + os.Getenv("APP_PORT") + "/")

	services.ConnectDatabse()

	route := mux.NewRouter()

	addApproutes(route)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), route))
}
