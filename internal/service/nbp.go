package service

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	DateLayout = "2006-01-02"
)

type ExchangeRatesDTO struct {
	Table    string
	Currency string
	Code     string
	Rates    []RateDTO
}

type RateDTO struct {
	No            string
	EffectiveDate string
	Mid           float32
	Bid           float32
	Ask           float32
}

type Rate struct {
	RateDTO
	Table    string
	Currency string
	Code     string
}

func (r *Rate) asTable() map[string]interface{} {
	out := map[string]interface{}{
		"table":         r.Table,
		"currency":      r.Currency,
		"code":          r.Code,
		"no":            r.No,
		"effectiveDate": r.EffectiveDate,
	}
	if r.Mid > 0 {
		out["mid"] = r.Mid
	} else {
		out["bid"] = r.Bid
		out["ask"] = r.Ask
	}
	return out
}

func (r *Rate) asJSON() interface{} {
	return r.asTable()
}

func GetRate(currency, table string, date time.Time) (*Rate, error) {
	/*
		Table A rates are published every day (working days only) between 11:45-12:15 CET
		Table B rates are published every Wednesday (if it's not working day then day before) between 11:45-12:15 CET
		Table C rates are published every day (working days only) between 7:45-8:15 CET

		To be sure we will get rates for any table we filter for 10 days period since requested date
		Rates are sorted by date in accending order, so the last rate is the closest to requested date
	*/
	dateFrom := date.AddDate(0, 0, -10)

	exRates, err := getExchangeRates(currency, table, dateFrom, date)
	if err != nil {
		return nil, err
	}

	latest := exRates.Rates[len(exRates.Rates)-1]

	rate := &Rate{
		latest,
		exRates.Table,
		exRates.Currency,
		exRates.Code,
	}

	return rate, nil
}

func getExchangeRates(currency, table string, dateFrom, dateTo time.Time) (*ExchangeRatesDTO, error) {
	client := resty.New().
		SetHostURL("https://api.nbp.pl").
		SetTimeout(1 * time.Second).
		SetRetryCount(3)

	var exRates ExchangeRatesDTO

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetPathParams(map[string]string{
			"table":    table,
			"currency": currency,
			"dateFrom": dateFrom.Format(DateLayout),
			"dateTo":   dateTo.Format(DateLayout),
		}).
		SetResult(&exRates).
		Get("/api/exchangerates/rates/{table}/{currency}/{dateFrom}/{dateTo}")

	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("NBP responded with status code: %d, body: %s", resp.StatusCode(), resp.Body())
	}

	return &exRates, nil
}
