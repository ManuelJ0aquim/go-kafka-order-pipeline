package usecases

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ManuelJ0aquim/go-kafka-order-pipeline/internal/domain"
)

type ProcessOrderUseCase struct{}

func NewProcessOrderUseCase() *ProcessOrderUseCase {
	return &ProcessOrderUseCase{}
}

func (uc *ProcessOrderUseCase) Execute(payload []byte) (*domain.Order, error) {
	var order domain.Order
	if err := json.Unmarshal(payload, &order); err != nil {
		return nil, fmt.Errorf("falha ao deserializar pedido: %w", err)
	}

	log.Printf("[USECASE] Processando pedido ID: %s | Cliente: %s | Valor: $%.2f", order.ID, order.CustomerID, order.Amount)

	// Regra de negócio simples / simulação de processamento
	order.Status = domain.StatusProcessed

	log.Printf("[USECASE] Pedido %s processado com sucesso! Status: %s", order.ID, order.Status)
	return &order, nil
}
