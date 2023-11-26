package web

import (
	"log"
	"net/http"

	"github.com/beschlz/dartsly-match-service/internal/core/matches"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	Port         string
	chiRouter    *chi.Mux
	matchService matches.MatchService
}

func (app *App) initRoutes(matchService matches.MatchService) {
	app.chiRouter.Post("/matches", app.PostMatches)
}

func NewApp(port string, matchService matches.MatchService) *App {
	app := &App{
		Port:      port,
		chiRouter: chi.NewRouter(),
	}

	app.chiRouter.Use(middleware.Logger)
	app.matchService = matchService
	app.initRoutes(matchService)

	return app
}

func (app *App) Run() {

	err := http.ListenAndServe(app.Port, app.chiRouter)

	if err != nil {
		log.Fatal(err)
	}
}
