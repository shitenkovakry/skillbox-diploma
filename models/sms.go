package models

type SMSDatum struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

type SMSData []*SMSDatum
