package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/FreshPrince99/Go_API/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter() // a struct for saving the api
	handlers.Handler(r)

	fmt.Println("Starting GO API service....")
	fmt.Println(`
	  /$$$$$$   /$$$$$$         /$$$$$$  /$$$$$$$  /$$$$$$
	/$$__  $$ /$$__  $$       /$$__  $$| $$__  $$|_  $$_/
	| $$  \__/| $$  \ $$      | $$  \ $$| $$  \ $$  | $$  
	| $$ /$$$$| $$  | $$      | $$$$$$$$| $$$$$$$/  | $$  
	| $$|_  $$| $$  | $$      | $$__  $$| $$____/   | $$  
	| $$  \ $$| $$  | $$      | $$  | $$| $$        | $$  
	|  $$$$$$/|  $$$$$$/      | $$  | $$| $$       /$$$$$$
	\______/  \______/       |__/  |__/|__/      |______/
                                                      
                                                      
                                                      `)
	
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Fatal(err)
	}
}