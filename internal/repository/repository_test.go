package repository_test

import (
	"testing"

	"github.com/SussyaPusya/TZ/internal/domain"
	"github.com/SussyaPusya/TZ/internal/repository"
)

func TestAddQuote(t *testing.T) {
	repo := repository.New()

	quote := domain.QuoteCell{
		Quote:  "Test quote",
		Author: "Author1",
	}

	id := repo.AddQuote(quote)

	if id != 1 {
		t.Errorf("expected id 1, got %d", id)
	}

	stored := repo.Storage[id]
	if stored.Quote != quote.Quote || stored.Author != quote.Author {
		t.Errorf("stored quote does not match original")
	}
}

func TestGetAllQuotes(t *testing.T) {
	repo := repository.New()

	q1 := domain.QuoteCell{Quote: "Quote 1", Author: "A1"}
	q2 := domain.QuoteCell{Quote: "Quote 2", Author: "A2"}

	id1 := repo.AddQuote(q1)
	id2 := repo.AddQuote(q2)

	all := repo.GetAllQuotes()

	if len(all) != 2 {
		t.Errorf("expected 2 quotes, got %d", len(all))
	}

	found := map[int]bool{id1: false, id2: false}
	for _, q := range all {
		found[q.ID] = true
	}

	for k, v := range found {
		if !v {
			t.Errorf("quote with ID %d not found", k)
		}
	}
}

func TestDeleteQuote(t *testing.T) {
	repo := repository.New()

	quote := domain.QuoteCell{Quote: "Delete me", Author: "Author"}
	id := repo.AddQuote(quote)

	repo.DeleteQoute(id)

	if _, exists := repo.Storage[id]; exists {
		t.Errorf("quote with id %d was not deleted", id)
	}
}

func TestGetQuotesFilterAuthor(t *testing.T) {
	repo := repository.New()

	q1 := domain.QuoteCell{Quote: "Q1", Author: "AuthorX"}
	q2 := domain.QuoteCell{Quote: "Q2", Author: "AuthorY"}
	q3 := domain.QuoteCell{Quote: "Q3", Author: "AuthorX"}

	repo.AddQuote(q1)
	repo.AddQuote(q2)
	repo.AddQuote(q3)

	result := repo.GetQuotesFilterAuthor("AuthorX")

	if len(*result) != 2 {
		t.Errorf("expected 2 quotes for AuthorX, got %d", len(*result))
	}
}
