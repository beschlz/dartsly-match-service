package web

import (
	"log"
	"net/http"

	"github.com/beschlz/dartsly-match-service/internal/core/matches"
	"github.com/beschlz/dartsly-match-service/internal/core/matches/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	Port      string
	chiRouter *chi.Mux
}

func (app *App) initRoutes(matchService matches.MatchService) {

	app.chiRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		match, err := (matchService.CreateMatch(domain.MatchCreationRequest{
			PlayerCount:      2,
			CheckOutSettings: "singleout",
		}))

		if err != nil {
			log.Fatal(err)
		}

		_, err = w.Write([]byte(match.CheckoutSettings))

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
