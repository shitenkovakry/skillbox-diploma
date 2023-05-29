package main

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

func main() {
	in := `RU;88;713;Topolo
	US;77;655;Rond
	GB;46;1144;Topolo
	FR;19;1668;Topolo
	BL;79;903;Kildy
	AT;0;237;Topolo
	BG;68;1004;Rond
	DK13;110;Topolo
	CA;81;1272;Rond
	ES;61;1838;Topolo
	CH;11;978;Topolo
	TR;12;857;Rond
	PE;37;735;Topolo
	NZ;57;1632;Kildy
	MC;82;919;Kildy
`
	// myfile, err := os.OpenFile("./data.sms", os.O_RDONLY, os.ModePerm)
	// if err != nil {
	// 	panic(err)
	// }

	buffer := strings.NewReader(in)
	csvReader := csv.NewReader(buffer)
	csvReader.Comma = ';'

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
			fmt.Println(err)
		}

		smsItems = append(smsItems, smsItem)
	}

	fmt.Println(smsItems)
}

var (
	approvedProviders = [3]string{"Topolo", "Rond", "Kildy"}
)

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
	for index := 0; index < len(approvedProviders); index++ {
		if strings.EqualFold(provider, approvedProviders[index]) {
			found = true

			break
		}
	}

	if !found {
		return nil, errors.New("provider not found in approved providers")
	}

	result := &models.SMSDatum{
		Country:      country.Alpha2(),
		Bandwidth:    bandwidth,
		ResponseTime: responseTime,
		Provider:     provider,
	}

	return result, nil
}
