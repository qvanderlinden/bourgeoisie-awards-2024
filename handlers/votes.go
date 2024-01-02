package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	myClerk "github.com/bourgeoisie-awards-2024/clerk"
	"github.com/bourgeoisie-awards-2024/models"
	"github.com/bourgeoisie-awards-2024/pages"
	"github.com/clerkinc/clerk-sdk-go/clerk"
)

type VotesPageHandler struct {
	clerkFrontendConfig *myClerk.FrontendConfig
	pool                *sql.DB
}

func (h VotesPageHandler) serveVotesPage(w http.ResponseWriter, r *http.Request) {
	claims, ok := clerk.SessionFromContext(r.Context())
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	m, err := models.NewVotesPageModel(h.pool, claims.Subject)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	component := pages.VotesPage(h.clerkFrontendConfig, m)
	component.Render(context.Background(), w)
}

func (h VotesPageHandler) handleVotes(r *http.Request) error {
	claims, ok := clerk.SessionFromContext(r.Context())
	if !ok {
		return fmt.Errorf("no clerk session found")
	}

	votes := map[string]string{}
	err := json.NewDecoder(r.Body).Decode(&votes)
	if err != nil {
		return err
	}

	tx, err := h.pool.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	for categoryId, nomineeId := range votes {
		_, err := tx.Exec(`
			INSERT INTO votes
				(user_id, category_id, nominee_id, updated_at)
			VALUES
				($1, $2, $3, CURRENT_TIMESTAMP)
			ON CONFLICT ON CONSTRAINT unique_vote_for_user_category 
			DO UPDATE SET
				nominee_id = $3,
				updated_at = CURRENT_TIMESTAMP
			`, claims.Subject, categoryId, nomineeId)

		if err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}

func (h VotesPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		h.serveVotesPage(w, r)
		return
	}
	if r.Method == "POST" {
		err := h.handleVotes(r)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func NewVotesPageHandler(clerkFrontendConfig *myClerk.FrontendConfig, pool *sql.DB) http.Handler {
	return VotesPageHandler{
		clerkFrontendConfig: clerkFrontendConfig,
		pool:                pool,
	}
}
