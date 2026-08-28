package main

import (
	"context"
	"log"
)

type ctxKey string

const requestIDKey ctxKey = "requestID"

func processRequest(ctx context.Context, requestID string) {
	ctx = context.WithValue(
		ctx,
		requestIDKey,
		requestID,
	)
	requestID_1 := ctx.Value(requestIDKey)
	log.Printf("processRequest %v", requestID_1)
	validateData(ctx)
}
func validateData(ctx context.Context) {
	requestID_2 := ctx.Value(requestIDKey)
	log.Printf("validateData %v", requestID_2)
	saveData(ctx)
}

func saveData(ctx context.Context) {
	requestID_3 := ctx.Value(requestIDKey)
	log.Printf("saveData %v", requestID_3)
}
func main() {
	processRequest(context.Background(), "abc-123")
}
