package voicecall_test

import (
	"os"
	voicecall "skillbox-diploma/datasources/voice-call"
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

	actual := voicecall.Parse(inputFile)
	expected := voicecall.Parse(expectedFile)

	assert.Equal(t, expected, actual)
}
