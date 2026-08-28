package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ManuelJ0aquim/go-kafka-order-pipeline/internal/domain"
	"github.com/ManuelJ0aquim/go-kafka-order-pipeline/internal/kafka"
)

func main() {
	brokers := []string{"localhost:9092"}
	topic := "orders-topic"

	producer := kafka.NewProducer(brokers, topic)
	defer producer.Close()

	log.Println("[PRODUCER] Iniciando envio de pedidos...")

	for i := 1; i <= 5; i++ {
		order := domain.Order{
			ID:         fmt.Sprintf("ord-%03d", i),
			CustomerID: fmt.Sprintf("cust-%03d", 100+i),
			Amount:     99.90 * float64(i),
			Status:     domain.StatusPending,
			CreatedAt:  time.Now(),
		}

		payload, err := json.Marshal(order)
		if err != nil {
			log.Printf("Erro ao serializar pedido: %v", err)
			continue
		}

		err = producer.Publish(context.Background(), []byte(order.ID), payload)
		if err != nil {
			log.Printf("Erro ao enviar mensagem: %v", err)
		} else {
			log.Printf("[PRODUCER] Publicado: Pedido %s | Valor: $%.2f", order.ID, order.Amount)
		}

		time.Sleep(1 * time.Second)
	}
}
