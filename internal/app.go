package internal

import (
	"database/sql"
	"fmt"
	"it-sloth/user.api/config"
	"it-sloth/user.api/internal/controller"
	"it-sloth/user.api/internal/factory"
	"it-sloth/user.api/internal/repository"
	"it-sloth/user.api/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type App struct {
	env *config.Env
	db *sql.DB
	publicController *controller.PublicController
	handler http.Handler
}

func (a *App) Build() error {
	connection, err := sql.Open("postgres", a.env.DbDsn)
	if err != nil {
		log.Fatal(err)
		return err
	}

	a.db = connection
	a.publicController = controller.NewPublicController(
		service.NewUser(
			repository.NewUserRepository(a.db),
			factory.NewEntityFactory(),
		),
	)

	muxHandler := a.buildRoutes()
	a.handler = cors.Default().Handler(muxHandler)

	return nil
}

func (a *App) buildRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user", a.publicController.Create).Methods("POST")

	return router
} 

func (a *App) Run() error {
	err := http.ListenAndServe(fmt.Sprintf(":%s", a.env.Port), a.handler)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (a *App) Shutdown() error {
	log.Fatal("shutting down")
	a.db.Close()

	return nil
}

func NewApp(env *config.Env) *App {
	return &App{
		env: env,
	}
}