package main

import (
	"github.com/bgoldovsky/teamer/auth-api/internal/cfg"
	"github.com/bgoldovsky/teamer/auth-api/internal/handlers"
)

func main() {
	secret := cfg.GetSecret()
	port := cfg.GetHTTPPort()

	app := handlers.New()
	app.Initialize(secret)
	app.Run(port)
}
