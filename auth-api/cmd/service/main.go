package main

import (
	"github.com/bgoldovsky/dutyer/auth-api/internal/cfg"
	"github.com/bgoldovsky/dutyer/auth-api/internal/handlers"
)

func main() {
	secret := cfg.GetSecret()
	port := cfg.GetHTTPPort()

	app := handlers.New()
	app.Initialize(secret)
	app.Run(port)
}
