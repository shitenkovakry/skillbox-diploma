package sms

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"skillbox-diploma/models"
	"strconv"
	"strings"

	"github.com/biter777/countries"
)

func Parse(data io.Reader) models.SMSData {
	csvReader := csv.NewReader(data)
	csvReader.Comma = ';'
	csvReader.FieldsPerRecord = -1

	var smsItems models.SMSData

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		smsItem, err := convertRecordToSMSDatum(record)
		if err != nil {
			log.Println(err)

			continue
		}

		smsItems = append(smsItems, smsItem)
	}

	return smsItems
}

func convertRecordToSMSDatum(record []string) (*models.SMSDatum, error) {
	if len(record) != 4 {
		return nil, errors.New("expected 4 items")
	}

	// obtain and validate country
	country := countries.ByName(record[0])
	if !country.IsValid() {
		return nil, errors.New("invalid country")
	}

	// obtain and validate bandwidth
	bandwidth, err := strconv.Atoi(record[1])
	if err != nil {
		return nil, err
	}
	if bandwidth < 0 || bandwidth > 100 {
		return nil, errors.New("incorrect bandwidth")
	}

	// obtain and validate responsetime
	responseTime, err := strconv.Atoi(record[2])
	if err != nil {
		return nil, err
	}
	if responseTime < 0 {
		return nil, errors.New("incorrect responseTime")
	}

	// obtain and validate provider
	provider := record[3]
	found := false
	for index := 0; index < len(models.ApprovedProviders); index++ {
		if strings.EqualFold(provider, models.ApprovedProviders[index]) {
			found = true

			break
		}
	}

	if !found {
		return nil, errors.New("provider not found in approved providers")
	}

	result := &models.SMSDatum{
		Country:      country.Alpha2(),
		Bandwidth:    fmt.Sprint(bandwidth),
		ResponseTime: fmt.Sprint(responseTime),
		Provider:     provider,
	}

	return result, nil
}
