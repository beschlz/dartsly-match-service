package web

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/beschlz/dartsly-match-service/internal/core/matches/domain"
)

const InternalServerError = "Internal server error"

func (app *App) PostMatches(w http.ResponseWriter, r *http.Request) {
	var requestBody MatchCreationRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	domainMatch := mapToDomainObject(requestBody)
	match, err := app.matchService.CreateMatch(domainMatch)

	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidMatchSetting):
			http.Error(w, InternalServerError, http.StatusBadRequest)
			return
		default:
			http.Error(w, InternalServerError, http.StatusInternalServerError)
			return
		}

	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(toDto(match)); err != nil {
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}

	if err != nil {
		log.Fatal(err)
	}
}
