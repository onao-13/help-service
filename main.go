package main

import (
	"context"
	"help-service/internals/config"
	"help-service/internals/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	cfg := config.UploadDevConfig()
	_, cancel := context.WithCancel(context.Background())

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)

	srv := server.New(cfg)

	go func() {
		osCall := <-s
		log.Println("Server stoped: ", osCall)
		srv.Shutdown()
		cancel()
	}()

	srv.Start()
}
