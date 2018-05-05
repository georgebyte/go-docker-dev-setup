package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setCors(w)
	fmt.Fprintf(w, "Hello, world!")
}

// used for COR preflight checks
func corsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setCors(w)
}

// util
func getClientUrl() string {
	if os.Getenv("APP_ENV") == "production" {
		return "http://localhost:3000" // change this to production domain
	} else {
		return "http://localhost:3000"
	}
}

func setCors(w http.ResponseWriter) {
	clientUrl := getClientUrl()
	w.Header().Set("Access-Control-Allow-Origin", clientUrl)
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.OPTIONS("/*any", corsHandler)

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running server in production mode")
	} else {
		log.Println("Running server in development mode")
	}

	http.ListenAndServe(":8080", router)
}
