package mms

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/biter777/countries"

	"skillbox-diploma/models"
)

type MMS struct {
	approvedProviders []string
	url               string
	client            http.Client
}

func New(url string) *MMS {
	// url = "http://localhost:8383/mms"
	return &MMS{
		approvedProviders: models.ApprovedProviders[:],
		url:               url,
		client:            http.Client{},
	}
}

func (mms *MMS) convertRecordToMMSDatum(record *models.MMSDatum) (*models.MMSDatum, error) {
	// obtain and validate country
	country := countries.ByName(record.Country)
	if !country.IsValid() {
		return nil, errors.New("invalid country")
	}

	// obtain and validate bandwidth
	bandwidth, err := strconv.Atoi(record.Bandwidth)
	if err != nil {
		return nil, err
	}
	if bandwidth < 0 || bandwidth > 100 {
		return nil, errors.New("incorrect bandwidth")
	}

	// obtain and validate responsetime
	responseTime, err := strconv.Atoi(record.ResponseTime)
	if err != nil {
		return nil, err
	}
	if responseTime < 0 {
		return nil, errors.New("incorrect responseTime")
	}

	// obtain and validate provider
	provider := record.Provider
	found := false
	for index := 0; index < len(mms.approvedProviders); index++ {
		if strings.EqualFold(provider, mms.approvedProviders[index]) {
			found = true

			break
		}
	}

	if !found {
		return nil, errors.New("provider not found in approved providers")
	}

	result := &models.MMSDatum{
		Country:      country.Alpha2(),
		Bandwidth:    fmt.Sprint(bandwidth),
		ResponseTime: fmt.Sprint(responseTime),
		Provider:     provider,
	}

	return result, nil
}

func (mms *MMS) Load() models.MMSData {
	response, err := mms.client.Get(mms.url)
	if err != nil {
		log.Print(err)

		return nil
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println("incorrect code")

		return nil
	}

	mmsData, err := mms.Parse(response.Body)

	if err != nil {
		log.Print(err)
	}

	return mmsData
}

func (mms *MMS) Parse(data io.Reader) (models.MMSData, error) {
	var (
		mmsDataRaw models.MMSData
		mmsData    models.MMSData
	)

	body, err := io.ReadAll(data)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &mmsDataRaw); err != nil {
		return nil, err
	}

	for i := 0; i < len(mmsDataRaw); i++ {
		record := mmsDataRaw[i]
		converted, err := mms.convertRecordToMMSDatum(record)
		if err != nil {
			log.Println(err)

			continue
		}

		mmsData = append(mmsData, converted)
	}

	return mmsData, nil
}
