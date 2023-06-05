package incident

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"skillbox-diploma/models"
	"strings"
)

type Incident struct {
	approvedStatusForIncident []string
	url                       string
	client                    http.Client
	data                      models.IncidentData
}

func New(url string) *Incident {
	// url = "http://localhost:8383/incident"
	incident := &Incident{
		approvedStatusForIncident: models.ApprovedStatusForIncident[:],
		url:                       url,
		client:                    http.Client{},
		data:                      nil,
	}

	incident.data = incident.Load()

	return incident
}

func (incident *Incident) Load() models.IncidentData {
	response, err := incident.client.Get(incident.url)
	if err != nil {
		log.Print(err)

		return nil
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println("incorrect code")

		return nil
	}

	incidentData, err := incident.Parse(response.Body)

	if err != nil {
		log.Print(err)
	}

	return incidentData
}

func (incident *Incident) Parse(data io.Reader) (models.IncidentData, error) {
	var (
		incidentDataRaw models.IncidentData
		incidentData    models.IncidentData
	)

	body, err := io.ReadAll(data)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &incidentDataRaw); err != nil {
		return nil, err
	}

	for index := 0; index < len(incidentDataRaw); index++ {
		record := incidentDataRaw[index]
		converted, err := incident.convertRecordToIncidentDatum(record)
		if err != nil {
			log.Println(err)

			continue
		}

		incidentData = append(incidentData, converted)
	}

	return incidentData, nil
}

func (incident *Incident) convertRecordToIncidentDatum(record *models.IncidentDatum) (*models.IncidentDatum, error) {
	// obtain and validate topic
	topic := record.Topic
	if topic == "" {
		return nil, errors.New("incorrect topic")
	}

	// obtain and validate status
	status := record.Status
	found := false
	for index := 0; index < len(incident.approvedStatusForIncident); index++ {
		if strings.EqualFold(status, incident.approvedStatusForIncident[index]) {
			found = true

			break
		}
	}

	if !found {
		return nil, errors.New("status not found in approved statuses")
	}

	result := &models.IncidentDatum{
		Topic:  topic,
		Status: status,
	}

	return result, nil
}

func (incident *Incident) Read() models.IncidentData {
	data := make(models.IncidentData, len(incident.data))

	for i := 0; i < len(data); i++ {
		data[i] = incident.data[i]
	}

	return data
}
