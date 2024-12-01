package main

import (
	"fmt"
	"it-sloth/user.api/config"
	"it-sloth/user.api/internal/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	publicController := controller.NewPublicController()

	router := mux.NewRouter()
	// router.HandleFunc("/user/{guid:.*}", publicController.Read).Methods("GET")
	router.HandleFunc("/user", publicController.Create).Methods("POST")

	return router
}
