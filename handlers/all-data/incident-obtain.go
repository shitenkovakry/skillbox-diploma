package alldata

import (
	"skillbox-diploma/models"
	"strings"
)

func (handler *Handler) obtainIncidentData() models.IncidentData {
	incidents := handler.incidents.Read()
	keyWord := "active"
	incidentsActive := models.IncidentData{}
	incidentsNotActive := models.IncidentData{}
	sorted := models.IncidentData{}

	for index := 0; index < len(incidents); index++ {
		elementOfIncident := incidents[index]

		if strings.EqualFold(elementOfIncident.Status, keyWord) {
			incidentsActive = append(incidentsActive, elementOfIncident)

			continue
		}

		incidentsNotActive = append(incidentsNotActive, elementOfIncident)
	}

	sorted = append(sorted, incidentsActive...)
	sorted = append(sorted, incidentsNotActive...)

	return sorted
}
