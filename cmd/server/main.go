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

	smsSource := sms.New("./skillbox/sms.data")
	mmsSource := mms.New("http://backend_skillbox:8383/mms")
	voiceCallSource := voicecall.New("./skillbox/voice.data")
	emailSource := email.New("./skillbox/email.data")
	billingSource := billing.New("./skillbox/billing.data")
	incidentSource := incident.New("http://backend_skillbox:8383/accendent")
	supportSource := support.New("http://backend_skillbox:8383/support")

	handlerForAllData := alldata.New(smsSource, mmsSource, voiceCallSource, emailSource, billingSource, incidentSource, supportSource)
	router.Method(http.MethodGet, "/all-data", handlerForAllData)

	router.HandleFunc("/chart.min.js",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./skillbox/chart.min.js")
		})

	router.HandleFunc("/main.css",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./skillbox/main.css")
		})

	router.HandleFunc("/main.js",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./skillbox/main.js")
		})

	router.HandleFunc("/status_page.html",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./skillbox/status_page.html")
		})

	router.HandleFunc("/true.png",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./skillbox/true.png")
		})

	router.HandleFunc("/false.png",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./skillbox/false.png")
		})

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
