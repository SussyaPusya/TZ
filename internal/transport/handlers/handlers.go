package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/SussyaPusya/TZ/internal/domain"
	"github.com/gorilla/mux"
)

type Repository interface {
	GetAllQuotes() []domain.QuoteCell

	GetQuotesFilterAuthor(author string) *[]domain.QuoteCell

	AddQuote(quote domain.QuoteCell) int
	DeleteQoute(id int)
}

type Handleres struct {
	repo Repository
}

func New(repo Repository) *Handleres {

	return &Handleres{repo: repo}
}

func (h *Handleres) GetQuotes(w http.ResponseWriter, r *http.Request) {

	author := r.URL.Query().Get("author")

	if author != "" {
		qoutes := h.repo.GetQuotesFilterAuthor(author)
		respJs, err := json.Marshal(qoutes)
		if err != nil {
			log.Print("filed to marshal json", err)

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJs)
		return

	}
	quotes := h.repo.GetAllQuotes()

	respJs, err := json.Marshal(quotes)
	if err != nil {
		log.Print("filed to marshal json", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJs)
}

func (h *Handleres) AddQuote(w http.ResponseWriter, r *http.Request) {
	bytew, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("filed to read body", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var qoute domain.QuoteCell

	err = json.Unmarshal(bytew, &qoute)
	if err != nil {
		log.Print("filed to unmarshal struct", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := h.repo.AddQuote(qoute)

	resp := map[string]int{"id": id}

	respJson, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)

}

func (h *Handleres) DeleteQoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	strID := vars["id"]

	id, err := strconv.Atoi(strID)
	if err != nil {

		log.Print("filed to convert id to int", err)

		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	h.repo.DeleteQoute(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	message := `{"message":"succesful"}`
	w.Write([]byte(message))

}
