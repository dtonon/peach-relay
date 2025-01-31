package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fiatjaf/eventstore/badger"
	"github.com/fiatjaf/khatru"
	"github.com/joho/godotenv"
	"github.com/nbd-wtf/go-nostr"
)

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

func main() {
	godotenv.Load(".env")

	relay := khatru.NewRelay()

	relay.Info.Name = "Peach Relay"
	relay.Info.PubKey = "xxxxxxxxxxdcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"
	relay.Info.Description = "A relay that manages NIP-69 orders"
	relay.Info.Icon = "https://....."

	db := badger.BadgerBackend{
		Path: getEnv("DB_PATH"),
	}
	if err := db.Init(); err != nil {
		panic(err)
	}

	relay.StoreEvent = append(relay.StoreEvent, db.SaveEvent)
	relay.QueryEvents = append(relay.QueryEvents, db.QueryEvents)
	relay.DeleteEvent = append(relay.DeleteEvent, db.DeleteEvent)
	relay.ReplaceEvent = append(relay.ReplaceEvent, db.ReplaceEvent)

	relay.RejectEvent = append(relay.RejectEvent, func(ctx context.Context, event *nostr.Event) (reject bool, msg string) {
		if event.Kind != 38383 {
			return true, "I accept only NIP-69 orders"
		}
		return false, ""
	})

	fmt.Println("🍑 Peach Relay running on :", getEnv("RELAY_PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", getEnv("RELAY_PORT")), relay)

}
