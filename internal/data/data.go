package data

import (
	"log"
	"strconv"
)

type AveragePrice struct {
	Date  string
	Price float64
}

type Results struct {
	Monthly       []float64      // 12 месяцев; кол-во потраченного
	Hourly        []float64      // 24 часа; частотное распределение
	TotalPaid     float64        // заплатили за весь период (gas price * value)
	AveragePrices []AveragePrice // средняя цена за каждый день
}

func NewResults(url string) (*Results, error) {
	gasPrice, err := downloadJson(url)
	if err != nil {
		return nil, err
	}

	monthly := make([]float64, 12)
	hourly := make([]float64, 24)
	hourlyCounts := make([]float64, 24)
	var total float64
	average := make([]AveragePrice, 0, len(gasPrice.Ethereum.Transactions)/24)
	var sum float64

	// ожидаем, что массив данных не имеет пропусков и отсортирован по времени
	for i, tx := range gasPrice.Ethereum.Transactions {
		month, err := strconv.Atoi(tx.Time[3:5])
		if err != nil {
			log.Printf("ERROR: %v", err)
			return nil, err
		}
		hour := i % 24

		monthly[month-1] += tx.GasValue

		hourlyCounts[hour]++
		hourly[hour] += (tx.GasPrice - hourly[hour]) / hourlyCounts[hour]

		total += tx.GasPrice * tx.GasValue

		sum += tx.GasPrice
		if hour == 23 {
			average = append(average, AveragePrice{tx.Time[:8], sum / 24})
			sum = 0
		}
	}

	return &Results{
		Monthly:       monthly,
		Hourly:        hourly,
		TotalPaid:     total,
		AveragePrices: average,
	}, nil
}
