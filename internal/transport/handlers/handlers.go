package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/SussyaPusya/TZ/internal/domain"
)

type Repository interface {
	GetAllQuotes() []domain.QuoteCell

	AddQuote(quote domain.QuoteCell) int
}

type Handleres struct {
	repo Repository
}

func New(repo Repository) *Handleres {

	return &Handleres{repo: repo}
}

func (h *Handleres) GetAllQuotes(w http.ResponseWriter, r *http.Request) {

	quotes := h.repo.GetAllQuotes()

	fmt.Fprintln(w, quotes)

	w.WriteHeader(http.StatusOK)
}

func (h *Handleres) AddQuote(w http.ResponseWriter, r *http.Request) {
	bytew, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("invalid json")
	}
	defer r.Body.Close()

	var qoute domain.QuoteCell

	err = json.Unmarshal(bytew, &qoute)
	if err != nil {
		log.Print("invalid json")
	}

	id := h.repo.AddQuote(qoute)

	resp := map[string]int{"id": id}

	respJson, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusOK)
}
