# Variáveis
APP_NAME := go-kafka-order-pipeline
DOCKER_COMPOSE := docker-compose.yml
PRODUCER_MAIN := cmd/producer/main.go
CONSUMER_MAIN := cmd/consumer/main.go

.PHONY: help up down restart status producer consumer build test fmt clean generator

## help: Exibe esta mensagem de ajuda
help:
	@echo "Uso: make [alvo]"
	@echo ""
	@echo "Alvos disponíveis:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'

## up: Sobe o container da infraestrutura (Kafka)
up:
	@echo "Iniciando Kafka..."
	docker compose up -d

## down: Para e remove o container da infraestrutura
down:
	@echo "Parando Kafka e Zookeeper..."
	docker compose down

## restart: Reinicia os serviços da infraestrutura
restart: down up

## status: Exibe o status dos containers em execução
status:
	docker compose ps

## producer: Executa a aplicação do Produtor
producer:
	@echo "Executando o Producer..."
	go run $(PRODUCER_MAIN)

## consumer: Executa a aplicação do Consumidor
consumer:
	@echo "Executando o Consumer..."
	go run $(CONSUMER_MAIN)

## build: Compila os binários do produtor e consumidor
build:
	@echo "Compilando os binários..."
	@mkdir -p bin
	go build -o bin/producer $(PRODUCER_MAIN)
	go build -o bin/consumer $(CONSUMER_MAIN)
	@echo "Binários salvos na pasta ./bin/"

## fmt: Formata o código fonte Go
fmt:
	go fmt ./...

## clean: Remove arquivos compilados e binários temporários
clean:
	@echo "Limpando binários..."
	rm -rf bin
