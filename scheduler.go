package main

import (
	"fmt"
	storage "github.com/iilinegor/test-taxi-service/storage"
	"time"
)

func runScheduler(ordersStorage *storage.Storage) {

	for len(ordersStorage.Orders) < storage.TokenCountLimit {
		time.Sleep(200 * time.Millisecond)
		ordersStorage.Shake()
	}
	fmt.Println("Token count limit reached")
}
