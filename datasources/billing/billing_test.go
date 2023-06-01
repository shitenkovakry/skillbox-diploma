package billing_test

import (
	"encoding/json"
	"os"
	"skillbox-diploma/datasources/billing"
	"skillbox-diploma/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	inputFile, err := os.OpenFile("./mocks/case1_input", os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	expectedFile, err := os.ReadFile("./mocks/case1_expected.json")
	if err != nil {
		t.Fatal(err)
	}

	actual := billing.Parse(inputFile)
	var expected *models.BillingDatum

	if err := json.Unmarshal(expectedFile, &expected); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
