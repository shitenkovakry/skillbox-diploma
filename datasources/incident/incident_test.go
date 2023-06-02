package incident_test

import (
	"encoding/json"
	"fmt"
	"os"
	"skillbox-diploma/datasources/incident"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_case1(t *testing.T) {
	inputFile, err := os.OpenFile("./mocks/case1_input.json", os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	expectedFile, err := os.OpenFile("./mocks/case1_expected.json", os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := incident.New("").Parse(inputFile)
	assert.Nil(t, err)

	expected, err := incident.New("").Parse(expectedFile)
	assert.Nil(t, err)

	dataToPrint, _ := json.Marshal(actual)
	fmt.Println(string(dataToPrint))

	assert.Equal(t, expected, actual)
}
