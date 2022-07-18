package main

import (
	"flag"
	"log"

	tgClient "github.com/EfimoffN/tgBotY/clients/telegram"
	"github.com/EfimoffN/tgBotY/consumer/eventconsumer"
	"github.com/EfimoffN/tgBotY/events/telegram"
	"github.com/EfimoffN/tgBotY/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := eventconsumer.New(eventProcessor, eventProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
