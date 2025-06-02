package repository

import (
	"sync"

	"github.com/SussyaPusya/TZ/internal/domain"
)

type Repository struct {
	Storage map[int]domain.QuoteCell
	nextId  int

	mu sync.Mutex
}

func New() *Repository {

	return &Repository{
		Storage: make(map[int]domain.QuoteCell),
		nextId:  1,
		mu:      sync.Mutex{},
	}
}

func (r *Repository) AddQuote(quote domain.QuoteCell) int {
	defer r.mu.Unlock()
	r.mu.Lock()
	id := r.nextId

	r.Storage[id] = quote

	r.nextId++

	return id

}

func (r *Repository) GetAllQuotes() []domain.QuoteCell {

	var res []domain.QuoteCell
	for id, val := range r.Storage {
		val.ID = id
		res = append(res, val)
	}

	return res
}

func (r *Repository) DeleteQoute(id int) {

	r.mu.Lock()

	delete(r.Storage, id)

	defer r.mu.Unlock()

}

func (r *Repository) GetQuotesFilterAuthor(author string) *[]domain.QuoteCell {
	r.mu.Lock()

	defer r.mu.Unlock()

	var result []domain.QuoteCell

	for id, val := range r.Storage {

		if val.Author == author {
			val.ID = id
			result = append(result, val)
		}

		continue
	}

	return &result
}
