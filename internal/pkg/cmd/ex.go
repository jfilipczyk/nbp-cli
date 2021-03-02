package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/gosuri/uitable"
	"github.com/jfilipczyk/nbp/internal/pkg/service"
	"github.com/spf13/cobra"
)

const isoLayout = "2006-01-02"

var (
	ExTable string
	ExDate  string

	exCmd = &cobra.Command{
		Use:   "ex [currency]",
		Short: "Get for exchange rate",
		Long:  "Get for exchange rate for given currency",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a currency argument")
			}
			if len(args[0]) != 3 {
				return fmt.Errorf("invalid currency given: %s", args[0])
			}
			if len(ExDate) > 0 {
				if _, err := time.Parse(isoLayout, ExDate); err != nil {
					return fmt.Errorf("exchange rate date has invalid format: %s", ExDate)
				}
			}
			if len(ExTable) != 1 {
				return fmt.Errorf("invalid table given: %s", ExTable)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			currency := args[0]
			var date time.Time
			if len(ExDate) > 0 {
				date, _ = time.Parse(isoLayout, ExDate)
			} else {
				date = time.Now()
			}

			rate, err := service.GetRate(currency, ExTable, date)
			if err != nil {
				return err
			}

			printRate(rate)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(exCmd)

	exCmd.Flags().StringVarP(&ExTable, "table", "t", "a", "Exchange rate table")
	exCmd.Flags().StringVarP(&ExDate, "date", "d", "", "Exchange rate date in ISO format YYYY-MM-DD")
}

func printRate(rate *service.Rate) {
	table := uitable.New()
	table.MaxColWidth = 80

	table.AddRow("No:", rate.No)
	table.AddRow("Date:", rate.EffectiveDate)
	if rate.Mid != 0 {
		table.AddRow("Mid:", rate.Mid)
	} else {
		table.AddRow("Bid:", rate.Bid)
		table.AddRow("Ask:", rate.Ask)
	}
	table.AddRow("") // blank

	fmt.Println(table)
}
