package dto

import (
	"time"

	"github.com/pablisson/go-gateway/internal/domain"
)

type CreateAccountInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateAccountResponse struct {
	ID    string `json:"id"`
	APIKey string `json:"api_key"`
}

type AccountOutput struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Balance float64 `json:"balance"`
	APIKey string `json:"api_key,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToAccount(input CreateAccountInput) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
}

func FromAccount(account *domain.Account) AccountOutput {
	return AccountOutput{
		ID: account.ID,
		Name: account.Name,
		Email: account.Email,
		Balance: account.Balance,
		APIKey: account.APIKey,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}