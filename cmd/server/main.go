package main

import (
	"log"
	"net/http"
	"skillbox-diploma/datasources/sms"
	alldata "skillbox-diploma/handlers/all-data"

	"github.com/go-chi/chi/v5"
)

const (
	addr = ":8383"
)

func main() {
	router := chi.NewRouter()

	smsSource := sms.New("./skillbox/sms.data")

	handlerForAllData := alldata.New(smsSource)
	router.Method(http.MethodGet, "/all/data", handlerForAllData)

	server := NewServer(addr, router)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server is error: %v", err)
	}
}

func NewServer(address string, router *chi.Mux) *http.Server {
	return &http.Server{
		Addr:    address,
		Handler: router,
	}
}
