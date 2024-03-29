package main

import (
	"github.com/beschlz/dartsly-match-service/internal/adapters/inmemory"
	"github.com/beschlz/dartsly-match-service/internal/adapters/web"
	"github.com/beschlz/dartsly-match-service/internal/core/matches"
)

func main() {

	matchRepo := inmemory.NewMatchRepo()
	matchService := matches.NewMatchService(matchRepo)

	app := web.NewApp(":8080", matchService)
	app.Run()
}
