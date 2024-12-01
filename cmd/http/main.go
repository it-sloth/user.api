package main

import (
	"it-sloth/user.api/config"
	"it-sloth/user.api/internal"
	"log"
)

func main() {
	app := internal.NewApp(config.GetEnv())
	err := app.Build()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer app.Shutdown()

	err = app.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}