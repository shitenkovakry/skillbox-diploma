package billing

import (
	"io"
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

	firstLine := allLines[0]
	encodedLine := encodeToLine([]byte(firstLine))
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
