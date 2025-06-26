package main

import (
	"acmmanager/internal/api/router"
	"crypto/tls"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	cert := "cert.pem"
	key := "key.pem"

	tlfConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	mux := router.Router()
	server := &http.Server{
		Addr:      os.Getenv("API_PORT"),
		Handler:   mux,
		TLSConfig: tlfConfig,
	}
	fmt.Println("server running on port 3000")
	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal(err)
	}
}
