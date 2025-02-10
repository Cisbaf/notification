package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"whatsappbot/src/server"
	"whatsappbot/src/wpp"
)

func main() {
	var wg sync.WaitGroup

	// Configura o contexto e o canal de sinalização
	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	wg.Add(1)
	go wpp.RunClient(ctx, &wg)
	srv := server.NewServer()
	go srv.Run(ctx, &wg)
	fmt.Println("Serviço iniciado. Pressione Ctrl+C para parar.")

	// Aguarda o sinal para finalizar
	<-sigChan
	fmt.Println("Sinal de encerramento recebido.")
	cancel() // Envia o cancelamento para as goroutines

	wg.Wait() // Aguarda as goroutines terminarem
	fmt.Println("Todas as goroutines foram finalizadas. Saindo do programa.")
}
