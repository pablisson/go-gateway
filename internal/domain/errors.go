package domain

import "errors"

var (
	// quando não encontrar uma conta
	ErrAccountNotFound = errors.New("account not found")

	// quando uma conta com a mesma chave de API já existe
	ErrDuplicateAPIKey = errors.New("duplicate api key")

	// quando não encontrar uma nota de fatura
	ErrInvoiceNoteFound = errors.New("invoice note not found")
	
	// quando o acesso não é autorizado
	ErrUnauthorizedAccess = errors.New("unauthorized access")
)