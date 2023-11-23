package main

import (
	"github.com/beschlz/dartsly-match-service/internal/adapters/persistance"
	"github.com/beschlz/dartsly-match-service/internal/adapters/web"
	matches "github.com/beschlz/dartsly-match-service/internal/core/services"
)

func main() {

	matchRepo := persistance.NewMatchRepo()
	matchService := matches.NewMatchService(matchRepo)

	app := web.NewApp(":8080", matchService)
	app.Run()
}
