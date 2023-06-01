package models

type EmailDatum struct {
	Country      string
	Provider     string
	DeliveryTime int
}

type EmailData []*EmailDatum
