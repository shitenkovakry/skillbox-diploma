package models

type SMSDatum struct {
	Country      string
	Bandwidth    int
	ResponseTime int
	Provider     string
}

type SMSData []*SMSDatum
