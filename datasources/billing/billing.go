package billing

import (
	"io"
	"log"
	"os"
	"skillbox-diploma/models"
	"strings"
)

func Parse(data io.Reader) *models.BillingDatum {
	allContent, err := io.ReadAll(data)
	if err != nil {
		return nil
	}

	allLines := strings.Split(string(allContent), "\n")
	if len(allLines) == 0 {
		return nil
	}

	firstLine := []byte(allLines[0])
	if len(firstLine) < 6 {
		return nil
	}

	encodedLine := encodeToLine(firstLine)
	result := &models.BillingDatum{}

	if encodedLine[0] == true {
		result.CreateCustomer = true
	}

	if encodedLine[1] == true {
		result.Purchase = true
	}

	if encodedLine[2] == true {
		result.Payout = true
	}

	if encodedLine[3] == true {
		result.Recurring = true
	}

	if encodedLine[4] == true {
		result.FraudControl = true
	}

	if encodedLine[5] == true {
		result.CheckoutPage = true
	}

	return result
}

func encodeToLine(line []byte) []bool {
	result := make([]bool, len(line))

	for i := 0; i < len(line); i++ {
		mask := string(line[i])

		if mask == "1" {
			result[len(line)-i-1] = true
		}
	}

	return result
}

type Billing struct {
	data *models.BillingDatum
}

func (billing *Billing) Read() *models.BillingDatum {
	return billing.data
}

func New(path string) *Billing {
	billingFile, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("cannot open file %s: %v", path, err)
	}

	data := Parse(billingFile)

	return &Billing{
		data: data,
	}
}
