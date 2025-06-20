package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)


type Account struct {
	ID string
	Name string
	Email string
	APIKey string
	Balance float64
	mu sync.RWMutex
	CreatedAt time.Time
	UpdatedAt time.Time
}



func NewAccount(name, email string) (*Account, error) {
	account := &Account{
		ID: uuid.New().String(),
		Name: name,
		Email: email,
		Balance: 0,
		APIKey: generateAPIKey(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return account, nil
}

func generateAPIKey() string {
	b := make([]byte, 16)
	rand.Read(b)

	return hex.EncodeToString(b)
}

func (a *Account) AddBalance(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
	a.UpdatedAt = time.Now()
}