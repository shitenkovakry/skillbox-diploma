package voicecall

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/biter777/countries"

	"skillbox-diploma/models"
)

func Parse(data io.Reader) models.VoiceCallData {
	csvReader := csv.NewReader(data)
	csvReader.Comma = ';'
	csvReader.FieldsPerRecord = -1

	var voiceCallItems models.VoiceCallData

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		voiceCallItem, err := convertRecordToVoiceCallDatum(record)
		if err != nil {
			log.Println(err)

			continue
		}

		voiceCallItems = append(voiceCallItems, voiceCallItem)
	}

	return voiceCallItems
}

func convertRecordToVoiceCallDatum(record []string) (*models.VoiceCallDatum, error) {
	if len(record) != 8 {
		return nil, errors.New("expected 8 items")
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
	for index := 0; index < len(models.ApprovedProvidersForVoiceCall); index++ {
		if strings.EqualFold(provider, models.ApprovedProvidersForVoiceCall[index]) {
			found = true

			break
		}
	}

	if !found {
		return nil, errors.New("provider not found in approved providers")
	}

	// obtain and validate connectionStability
	connectionStabilityInFloat64, err := strconv.ParseFloat(record[4], 32)
	if err != nil {
		return nil, err
	}

	connectionStability := float32(connectionStabilityInFloat64)
	if connectionStability < 0 {
		return nil, errors.New("incorrect connectionStability")
	}

	// obtain and validate TTFB
	ttfb, err := strconv.Atoi(record[5])
	if err != nil {
		return nil, err
	}
	if ttfb < 0 {
		return nil, errors.New("incorrect TTFB")
	}

	// obtain and validate voicePurity
	voicePurity, err := strconv.Atoi(record[6])
	if err != nil {
		return nil, err
	}
	if voicePurity < 0 {
		return nil, errors.New("incorrect voicePurity")
	}

	// obtain and validate medianOfCallsTime
	medianOfCallsTime, err := strconv.Atoi(record[7])
	if err != nil {
		return nil, err
	}
	if medianOfCallsTime < 0 {
		return nil, errors.New("incorrect medianOfCallsTime")
	}

	result := &models.VoiceCallDatum{
		Country:             country.Alpha2(),
		Bandwidth:           fmt.Sprint(bandwidth),
		ResponseTime:        fmt.Sprint(responseTime),
		Provider:            provider,
		ConnectionStability: connectionStability,
		TTFB:                ttfb,
		VoicePurity:         voicePurity,
		MedianOfCallsTime:   medianOfCallsTime,
	}

	return result, nil
}
