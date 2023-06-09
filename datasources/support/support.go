package support

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"skillbox-diploma/models"
	"strings"

	"github.com/pkg/errors"
)

type Support struct {
	approvedTopics []string
	url            string
	client         http.Client
	data           models.SupportData
}

func New(url string) *Support {
	// url = "http://localhost:8383/support"
	support := &Support{
		approvedTopics: models.ApprovedTopics[:],
		url:            url,
		client:         http.Client{},
		data:           nil,
	}

	support.data = support.Load()

	return support
}

func (support *Support) Load() models.SupportData {
	response, err := support.client.Get(support.url)
	if err != nil {
		log.Print(err)

		return nil
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println("incorrect code")

		return nil
	}

	supportData, err := support.Parse(response.Body)

	if err != nil {
		log.Print(err)
	}

	return supportData
}

func (support *Support) Parse(data io.Reader) (models.SupportData, error) {
	var (
		supportDataRaw models.SupportData
		supportData    models.SupportData
	)

	body, err := io.ReadAll(data)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &supportDataRaw); err != nil {
		return nil, err
	}

	for index := 0; index < len(supportDataRaw); index++ {
		record := supportDataRaw[index]
		converted, err := support.convertRecordToSupportDatum(record)
		if err != nil {
			log.Println(err)

			continue
		}

		supportData = append(supportData, converted)
	}

	return supportData, nil
}

func (support *Support) convertRecordToSupportDatum(record *models.SupportDatum) (*models.SupportDatum, error) {
	// obtain and validate activeTicket
	activeTickets := record.ActiveTickets
	if activeTickets < 0 {
		return nil, errors.New("incorrect activeTickets")
	}

	// obtain and validate topic
	topic := record.Topic
	found := false
	for index := 0; index < len(support.approvedTopics); index++ {
		if strings.EqualFold(topic, support.approvedTopics[index]) {
			found = true

			break
		}
	}

	if !found {
		return nil, errors.Wrapf(errors.New("topic not found in approved topics"), "topic = %s", topic)
	}

	result := &models.SupportDatum{
		Topic:         topic,
		ActiveTickets: activeTickets,
	}

	return result, nil
}

func (support *Support) Read() models.SupportData {
	data := make(models.SupportData, len(support.data))

	for i := 0; i < len(data); i++ {
		data[i] = support.data[i]
	}

	return data
}
