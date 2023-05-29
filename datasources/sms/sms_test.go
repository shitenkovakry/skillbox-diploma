package sms_test

import (
	"os"
	"skillbox-diploma/datasources/sms"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	inputFile, err := os.OpenFile("./mocks/case1_input.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	expectedFile, err := os.OpenFile("./mocks/case1_expected.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	actual := sms.Parse(inputFile)
	expected := sms.Parse(expectedFile)

	assert.Equal(t, expected, actual)
}
