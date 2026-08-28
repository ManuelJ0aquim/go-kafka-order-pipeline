package domain

import "time"

type OrderStatus string

const (
	StatusPending   OrderStatus = "PENDING"
	StatusProcessed OrderStatus = "PROCESSED"
	StatusFailed    OrderStatus = "FAILED"
)

type Order struct {
	ID         string      `json:"id"`
	CustomerID string      `json:"customer_id"`
	Amount     float64     `json:"amount"`
	Status     OrderStatus `json:"status"`
	CreatedAt  time.Time   `json:"created_at"`
}
