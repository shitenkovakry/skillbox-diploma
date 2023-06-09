package main

import (
	"log"
	"net/http"
	"skillbox-diploma/datasources/billing"
	"skillbox-diploma/datasources/email"
	"skillbox-diploma/datasources/incident"
	"skillbox-diploma/datasources/mms"
	"skillbox-diploma/datasources/sms"
	"skillbox-diploma/datasources/support"
	voicecall "skillbox-diploma/datasources/voice-call"
	alldata "skillbox-diploma/handlers/all-data"

	"github.com/go-chi/chi/v5"
)

const (
	addr = ":8080"
)

func main() {
	router := chi.NewRouter()

	smsSource := sms.New("../../skillbox/sms.data")
	mmsSource := mms.New("http://localhost:8383/mms")
	voiceCallSource := voicecall.New("../../skillbox/voice.data")
	emailSource := email.New("../../skillbox/email.data")
	billingSource := billing.New("../../skillbox/billing.data")
	incidentSource := incident.New("http://localhost:8383/incident")
	supportSource := support.New("http://localhost:8383/support")

	handlerForAllData := alldata.New(smsSource, mmsSource, voiceCallSource, emailSource, billingSource, incidentSource, supportSource)
	router.Method(http.MethodGet, "/all-data", handlerForAllData)

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
