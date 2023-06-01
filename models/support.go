package models

type SupportDatum struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

type SupportData []*SupportDatum
