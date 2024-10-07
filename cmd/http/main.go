package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"it-sloth/user.api/config"
	"it-sloth/user.api/internal/handler"
	"log"
	"net/http"
)

func main() {
	env := config.GetEnv()
	muxHandler := initRoutes()

	corsHandler := cors.Default().Handler(muxHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%s", env.Port), corsHandler)
	if err != nil {
		log.Fatal(err)
	}
}

func initRoutes() *mux.Router {
	publicHandler := handler.NewUser()

	router := mux.NewRouter()
	router.HandleFunc("/user/{guid:.*}", publicHandler.Read).Methods("GET")

	return router
}
