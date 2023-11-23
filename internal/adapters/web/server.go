package web

import (
	"net/http"

	matches "github.com/beschlz/dartsly-match-service/internal/core/ports"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	Port      string
	chiRouter *chi.Mux
}

func (app *App) initRoutes(matchService matches.MatchService) {
	app.chiRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(matchService.CreateMatch()))
	})
}

func NewApp(port string, matchService matches.MatchService) *App {
	app := &App{
		Port:      port,
		chiRouter: chi.NewRouter(),
	}

	app.chiRouter.Use(middleware.Logger)
	app.initRoutes(matchService)

	return app
}

func (app *App) Run() {

	http.ListenAndServe(app.Port, app.chiRouter)
}
