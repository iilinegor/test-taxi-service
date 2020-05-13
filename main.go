package main

import (
	storage "github.com/iilinegor/test-taxi-service/storage"
)

func main() {
	// init storage
	ordersStorage := storage.InitStorage(50)

	// start storage updater
	go runScheduler(&ordersStorage)

	// start API
	router(&ordersStorage)
}
