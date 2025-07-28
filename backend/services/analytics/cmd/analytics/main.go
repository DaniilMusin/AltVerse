package main

import (
    "context"
    "log"
    "os"

    "github.com/segmentio/kafka-go"
)

func envOr(k, d string) string {
    if v := os.Getenv(k); v != "" { return v }
    return d
}

func ProcessEvent(msg kafka.Message) {}

func main() {
    ctx := context.Background()
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{envOr("KAFKA_BROKERS", "kafka:9092")},
        Topic:   "gameplay.events.all",
        GroupID: "analytics",
    })
    defer r.Close()

    for {
        msg, err := r.ReadMessage(ctx)
        if err != nil { log.Fatal(err) }

        ProcessEvent(msg)
    }
}
