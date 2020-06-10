package main

import (
	"log"
	"os"
	"tg-worthlisteningto/pkg/client"
	"tg-worthlisteningto/pkg/telegram"
)

func main() {
	an := getEnvOrErr("AZ_ACCOUNT_NAME")
	at := getEnvOrErr("AZ_ACCOUNT_TOKEN")
	tg := getEnvOrErr("TELEGRAM_TOKEN")

	cl := client.Azure{}
	err := cl.Init(an, at)
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}

	_, updates, err := telegram.BotInit(tg)
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}

	err = telegram.Run(updates, cl)
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
}

func getEnvOrErr(k string) string {
	v, success := os.LookupEnv(k)
	if !success {
		log.Fatalf("Env variable %s cannot be empty", k)
	}
	return v
}
