package models

type IncidentDatum struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

type IncidentData []*IncidentDatum
