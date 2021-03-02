package cmd

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/jfilipczyk/nbp-cli/internal/service"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(exCmd)

	exCmd.Flags().StringVarP(&exTable, "table", "t", "a", "exchange rate table, one of: a, b, c")
	exCmd.Flags().StringVarP(&exDate, "date", "d", "", "exchange rate date in format YYYY-MM-DD")
}

var (
	exTable          string
	exDate           string
	exCurrencyRegexp = regexp.MustCompile("(?i)^[a-z]{3}$")
	exTableRegexp    = regexp.MustCompile("(?i)^[abc]{1}$")

	exCmd = &cobra.Command{
		Use:   "ex [currency]",
		Short: "Get exchange rate",
		Long:  "Get exchange rate for given currency",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("currency argument is required")
			}
			if !exCurrencyRegexp.MatchString(args[0]) {
				return fmt.Errorf("currency should be 3 alpha chars e.g. usd, given: %s", args[0])
			}
			if _, err := parseDateOrNow(exDate); err != nil {
				return fmt.Errorf("exchange rate date should have format YYYY-MM-DD, given: %s", exDate)
			}
			if !exTableRegexp.MatchString(exTable) {
				return fmt.Errorf("exchange rate table should be one of [a,b,c], given: %s", exTable)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			currency := args[0]
			date, _ := parseDateOrNow(exDate) // error check done in args validation

			rate, err := service.GetRate(currency, exTable, date)
			if err != nil {
				return err
			}

			return service.Output(outputFormat, rate)
		},
	}
)

func parseDateOrNow(exDate string) (date time.Time, err error) {
	date = time.Now()
	if len(exDate) == 0 {
		return
	}
	date, err = time.Parse(service.DateLayout, exDate)
	return
}
