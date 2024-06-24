package service

type ReservationRequest struct {
	EventID    string   `json:"event_id"`
	Spots      []string `json:"spots"`
<<<<<<< HEAD
	TicketType string   `json:"ticket_type"`
=======
	TicketKind string   `json:"ticket_kind"`
>>>>>>> 3febb45 (Mudanças nos docker-compose das pastas referentes ao nest, next e golang para rodar o projeto como um todo. Criação do readme.md final para melhor compreensão da execução do projeto.)
	CardHash   string   `json:"card_hash"`
	Email      string   `json:"email"`
}

type ReservationResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Spot       string `json:"spot"`
<<<<<<< HEAD
	TicketType string `json:"ticket_kind"`
=======
	TicketKind string `json:"ticket_kind"`
>>>>>>> 3febb45 (Mudanças nos docker-compose das pastas referentes ao nest, next e golang para rodar o projeto como um todo. Criação do readme.md final para melhor compreensão da execução do projeto.)
	Status     string `json:"status"`
	EventID    string `json:"event_id"`
}

type Partner interface {
	MakeReservation(req *ReservationRequest) ([]ReservationResponse, error)
}
