package main

import (
	"fmt"
	"github.com/iilinegor/test-taxi-service/storage"
	"log"
	"net/http"
)

func router(ordersStorage *storage.Storage) {
	fmt.Println("started on :8080")

	requestHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, ordersStorage.GetOrder())
	}

	adminRequestHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s", ordersStorage.GetStats())
	}

	http.HandleFunc("/request", requestHandler)
	http.HandleFunc("/admin/request", adminRequestHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
