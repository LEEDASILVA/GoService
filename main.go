package main

import (
	"log"
	"net/http"
	"os"

	server "./server"
	"github.com/joho/godotenv"
)

// create some env variables for the server addr and so on...
var (
	Cert_File   string
	Key_File    string
	Server_Addr string
)

func main() {
	// inject the logger
	// os.Stdout for locations where to print
	// and the flags for the standard flags and the filename
	logger := log.New(os.Stdout, "log ", log.LstdFlags|log.Lshortfile)

	// dependency injection
	h := server.InitHandler(logger)

	// load env file
	if err := godotenv.Load(".env"); err != nil {
		logger.Fatal(err)
	}
	Cert_File = os.Getenv("CERT_FILE")
	Key_File = os.Getenv("KEY_FILE")
	Server_Addr = os.Getenv("SERVER_ADDR")
	// takes the request and rotes them
	// it returns a serverMux structure
	// ServeMux is an HTTP request multiplexer.
	// It matches the URL of each incoming request against a list of registered
	// patterns and calls the handler for the pattern that
	// most closely matches the URL.
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.Logger(h.Home))

	server := server.NewSecurServer(mux, Server_Addr)

	logger.Println("server started")
	// because we are not using the default http (TLS for https)
	if err := server.ListenAndServeTLS(Cert_File, Key_File); err != nil {
		logger.Fatal(err)
	}
}
