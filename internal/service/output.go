package service

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/gosuri/uitable"
)

type Outputable interface {
	asJSON() interface{}
	asTable() map[string]interface{}
}

func Output(format string, v Outputable) error {
	if format == "json" {
		return outputJSON(v)
	}
	return outputTable(v)
}

func outputJSON(v Outputable) error {
	data, err := json.Marshal(v.asJSON())
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func outputTable(v Outputable) error {
	data := v.asTable()
	keys := getMapKeys(data)
	sort.Strings(keys)

	table := uitable.New()
	for _, k := range keys {
		table.AddRow(k, data[k])
	}
	table.AddRow("") // blank

	fmt.Println(table)
	return nil
}

func getMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
