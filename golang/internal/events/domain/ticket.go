package domain

import (
	"errors"

	"github.com/google/uuid"
)

// Errors
var (
<<<<<<< HEAD
	ErrInvalidTicketType = errors.New("invalid ticket type")
)

// TicketType represents the type of a ticket.
type TicketType string

const (
	TicketTypeHalf TicketType = "half" // Half-price ticket
	TicketTypeFull TicketType = "full" // Full-price ticket
)

// IsValidTicketType checks if a ticket type is valid.
func IsValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
=======
	ErrInvalidTicketKind = errors.New("invalid ticket kind")
)

// TicketKind represents the kind of a ticket.
type TicketKind string

const (
	TicketKindHalf TicketKind = "half" // Half-price ticket
	TicketKindFull TicketKind = "full" // Full-price ticket
)

// IsValidTicketKind checks if a ticket kind is valid.
func IsValidTicketKind(ticketKind TicketKind) bool {
	return ticketKind == TicketKindHalf || ticketKind == TicketKindFull
>>>>>>> 3febb45 (Mudanças nos docker-compose das pastas referentes ao nest, next e golang para rodar o projeto como um todo. Criação do readme.md final para melhor compreensão da execução do projeto.)
}

// Ticket represents a ticket for an event.
type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
<<<<<<< HEAD
	TicketType TicketType
=======
	TicketKind TicketKind
>>>>>>> 3febb45 (Mudanças nos docker-compose das pastas referentes ao nest, next e golang para rodar o projeto como um todo. Criação do readme.md final para melhor compreensão da execução do projeto.)
	Price      float64
}

// NewTicket creates a new ticket with the given parameters.
<<<<<<< HEAD
func NewTicket(event *Event, spot *Spot, ticketType TicketType) (*Ticket, error) {
	if !IsValidTicketType(ticketType) {
		return nil, ErrInvalidTicketType
=======
func NewTicket(event *Event, spot *Spot, ticketKind TicketKind) (*Ticket, error) {
	if !IsValidTicketKind(ticketKind) {
		return nil, ErrInvalidTicketKind
>>>>>>> 3febb45 (Mudanças nos docker-compose das pastas referentes ao nest, next e golang para rodar o projeto como um todo. Criação do readme.md final para melhor compreensão da execução do projeto.)
	}

	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
<<<<<<< HEAD
		TicketType: ticketType,
=======
		TicketKind: ticketKind,
>>>>>>> 3febb45 (Mudanças nos docker-compose das pastas referentes ao nest, next e golang para rodar o projeto como um todo. Criação do readme.md final para melhor compreensão da execução do projeto.)
		Price:      event.Price,
	}
	ticket.CalculatePrice()
	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
}

<<<<<<< HEAD
// CalculatePrice calculates the price based on the ticket type.
func (t *Ticket) CalculatePrice() {
	if t.TicketType == TicketTypeHalf {
=======
// CalculatePrice calculates the price based on the ticket kind.
func (t *Ticket) CalculatePrice() {
	if t.TicketKind == TicketKindHalf {
>>>>>>> 3febb45 (Mudanças nos docker-compose das pastas referentes ao nest, next e golang para rodar o projeto como um todo. Criação do readme.md final para melhor compreensão da execução do projeto.)
		t.Price /= 2
	}
}

// Validate checks if the ticket data is valid.
func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return errors.New("ticket price must be greater than zero")
	}
	return nil
}
