package tests

import (
	"AppMagicTestTask/internal/data"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var jsonFile = "test.json"
var testUrl = "https://raw.githubusercontent.com/CryptoRStar/GasPriceTestTask/main/gas_price.json"

func TestJSON(t *testing.T) {
	file, err := os.ReadFile(jsonFile)
	if err != nil {
		assert.Error(t, err)
	}

	var testResults data.Results

	err = json.Unmarshal(file, &testResults)
	if err != nil {
		assert.Error(t, err)
	}

	results, err := data.NewResults(testUrl)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, testResults, *results)
}
