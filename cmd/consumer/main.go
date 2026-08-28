package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ManuelJ0aquim/go-kafka-order-pipeline/internal/kafka"
	"github.com/ManuelJ0aquim/go-kafka-order-pipeline/internal/usecases"
)

func main() {
	brokers := []string{"localhost:9092"}
	topic := "orders-topic"
	groupID := "order-processing-group"

	consumer := kafka.NewConsumer(brokers, topic, groupID)
	defer consumer.Close()

	processUseCase := usecases.NewProcessOrderUseCase()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("[CONSUMER] Encerrando consumidor...")
		cancel()
	}()

	log.Println("[CONSUMER] Consumidor pronto. Aguardando mensagens...")

	for {
		// 1. Busca a mensagem do Kafka
		msg, err := consumer.FetchMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				break // Aplicação foi encerrada pelo sinal (Ctrl+C)
			}
			log.Printf("Erro ao buscar mensagem: %v", err)
			continue
		}

		log.Printf("[CONSUMER] Mensagem recebida | Partição: %d | Offset: %d | Key: %s",
			msg.Partition, msg.Offset, string(msg.Key))

		// 2. Processa o pedido
		_, err = processUseCase.Execute(msg.Value)
		if err != nil {
			log.Printf("Erro ao processar pedido: %v", err)
			continue
		}

		// 3. Confirma o offset apenas após o sucesso
		if err := consumer.CommitMessage(ctx, msg); err != nil {
			log.Printf("Erro ao fazer commit do offset: %v", err)
		}
	}
}
