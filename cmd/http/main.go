package main

import (
	"fmt"
	"github.com/rs/cors"
	"it-sloth/user.api/config"
	"it-sloth/user.api/internal/handler"
	"log"
	"net/http"
)

func main() {
	env := config.GetEnv()
	mux := initRoutes()

	corsHandler := cors.Default().Handler(mux)
	err := http.ListenAndServe(fmt.Sprintf(":%s", env.Port), corsHandler)
	if err != nil {
		log.Fatal(err)
	}
}

func initRoutes() *http.ServeMux {
	publicHandler := handler.NewPublic()

	mux := http.NewServeMux()
	mux.HandleFunc("/user", publicHandler.GetUser)

	return mux
}
