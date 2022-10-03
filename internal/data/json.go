package data

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Tx struct {
	Time           string
	GasPrice       float64
	GasValue       float64
	Average        float64
	MaxGasPrice    float64
	MedianGasPrice float64
}

type Ethereum struct {
	Transactions []Tx
}

type GasPriceJson struct {
	Ethereum Ethereum
}

func downloadJson(url string) (*GasPriceJson, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}

	var gasPrice GasPriceJson
	err = json.Unmarshal(respBytes, &gasPrice)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}

	return &gasPrice, nil
}
