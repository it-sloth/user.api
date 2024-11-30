package main

import (
	"fmt"
	"it-sloth/user.api/config"
	"it-sloth/user.api/internal/convertor"
	"it-sloth/user.api/internal/handler"
	"it-sloth/user.api/internal/repository"
	"it-sloth/user.api/internal/service"
	"it-sloth/user.api/internal/wrapper"
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
	publicHandler := handler.NewUser(
		service.NewUserService(
			repository.NewUserRepository(config.GetEnv()),
			convertor.NewUserEntityConvertor(),
		),
		wrapper.NewResponseWriter(),
	)

	router := mux.NewRouter()
	router.HandleFunc("/user/{guid:.*}", publicHandler.Read).Methods("GET")

	return router
}
