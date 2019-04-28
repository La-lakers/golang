// alarmDealList.go
package main

import (
	"fmt"

	"github.com/Luxurioust/excelize"
)

var filename = "./alarm_data_20190417_20190423aa.xlsx"

func readXlsxFile(filename string) {
	result := make(map[string]map[string]int)
	xlxs, err := excelize.OpenFile(filename)
	if err != nil {
		panic(err)
	}
	rows, err := xlxs.GetRows("alarm_deal_list")
	if err != nil {
		panic(err)
	}
	for _, row := range rows[1:] {
		item := row[2]
		content := row[3]
		if m, ok := result[item]; ok {
			m[content]++
		} else {
			m = make(map[string]int)
			m[content] = 1
			result[item] = m
		}
		//date := row[5]
	}
	for item, i := range result {
		for k, v := range i {
			fmt.Printf("Item:%s\tContent:%s\tCount:%d\n", item, k, v)
		}
	}

}

func main() {
	readXlsxFile(filename)
}
