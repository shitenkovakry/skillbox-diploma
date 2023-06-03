package email

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"skillbox-diploma/models"
	"strconv"
	"strings"

	"github.com/biter777/countries"
)

func Parse(data io.Reader) models.EmailData {
	csvReader := csv.NewReader(data)
	csvReader.Comma = ';'
	csvReader.FieldsPerRecord = -1

	var emailItems models.EmailData

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		emailItem, err := convertRecordToEmailDatum(record)
		if err != nil {
			log.Println(err)

			continue
		}

		emailItems = append(emailItems, emailItem)
	}

	return emailItems
}

func convertRecordToEmailDatum(record []string) (*models.EmailDatum, error) {
	if len(record) != 3 {
		return nil, errors.New("expected 3 items")
	}

	// obtain and validate country
	country := countries.ByName(record[0])
	if !country.IsValid() {
		return nil, errors.New("invalid country")
	}

	// obtain and validate provider
	provider := record[1]
	found := false
	for index := 0; index < len(models.ApprovedProvidersForEmail); index++ {
		if strings.EqualFold(provider, models.ApprovedProvidersForEmail[index]) {
			found = true

			break
		}
	}

	if !found {
		return nil, errors.New("provider not found in approved providers")
	}

	// obtain and validate deliveryTime
	deliveryTime, err := strconv.Atoi(record[2])
	if err != nil {
		return nil, err
	}
	if deliveryTime < 0 {
		return nil, errors.New("incorrect deliveryTime")
	}

	result := &models.EmailDatum{
		Country:      country.Alpha2(),
		Provider:     provider,
		DeliveryTime: deliveryTime,
	}

	return result, nil
}

type Email struct {
	data models.EmailData
}

func (email *Email) Read() models.EmailData {
	data := make(models.EmailData, len(email.data))

	for i := 0; i < len(data); i++ {
		data[i] = email.data[i]
	}

	return data
}

func New(path string) *Email {
	emailFile, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("cannot open file %s: %v", path, err)
	}

	data := Parse(emailFile)

	return &Email{
		data: data,
	}
}
