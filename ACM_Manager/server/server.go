package main

import (
	"acmmanager/internal/api/router"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {

	//err := godotenv.Load()
	//if err != nil {
	//	panic(err)
	//}

	mux := router.Router()
	server := &http.Server{
		Addr:    os.Getenv("API_PORT"),
		Handler: mux,
	}
	log.Println("server running on port 3000")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
