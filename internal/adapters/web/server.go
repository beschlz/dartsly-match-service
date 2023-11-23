package web

import (
	"log"
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
		_, err := w.Write([]byte(matchService.CreateMatch()))

		if err != nil {
			log.Fatal(err)
		}
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

	err := http.ListenAndServe(app.Port, app.chiRouter)

	if err != nil {
		log.Fatal(err)
	}
}
