package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	baseUrl   = "https://api.nbp.pl/api/exchangerates/rates"
	isoLayout = "2006-01-02"
)

type ExchangeRates struct {
	Table    string
	Currency string
	Code     string
	Rates    []Rate
}

type Rate struct {
	No            string
	EffectiveDate string
	Mid           float32
	Bid           float32
	Ask           float32
}

func GetRate(currency, table string, date time.Time) (*Rate, error) {
	dateFrom := date.AddDate(0, 0, -7) // rates are published only on working days

	exRates, err := getExchangeRates(currency, table, dateFrom, date)
	if err != nil {
		return nil, err
	}

	latest := exRates.Rates[len(exRates.Rates)-1]

	return &latest, nil
}

func getExchangeRates(currency, table string, dateFrom, dateTo time.Time) (*ExchangeRates, error) {
	dateFromStr := dateFrom.Format(isoLayout)
	dateToStr := dateTo.Format(isoLayout)

	url := fmt.Sprintf(baseUrl+"/%v/%v/%v/%v?format=json", table, currency, dateFromStr, dateToStr)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("NBP responded with status code: %d, body: %s", resp.StatusCode, body)
	}

	var exRates ExchangeRates
	err = json.Unmarshal(body, &exRates)
	if err != nil {
		return nil, err
	}

	return &exRates, nil
}
