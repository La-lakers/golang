// alarmAnalysis project main.go
package main

import (
	//"alarmAnalysis/models"
	"fmt"
	"regexp"
	"strings"

	"github.com/Luxurioust/excelize"
)

const sep = "<hupu>"

var rawAmarm = make(map[string]int)
var alarmDeal = make(map[string]int)

func AnalysisAlarm(rows [][]string, rawAmarm *map[string]int) {
	//result := make(map[string]int)
	nodeList := `note\[[^]]+]`
	nodeRe := regexp.MustCompile(nodeList)
	for _, row := range rows {
		if len(row) < 6 {
			continue
		}
		note := nodeRe.FindString(row[3])
		if note == "" {
			note = row[2]
		}
		date := strings.Split(row[5], " ")[0]
		key := row[1] + sep + row[2] + sep + note + sep + date
		(*rawAmarm)[key]++
	}
	//return result
}

func analysisDeal(rows [][]string, alarmDeal *map[string]int) {
	for _, row := range rows[1:] {
		if len(row) < 4 {
			continue
		}
		key := row[1] + sep + row[2] + sep + row[3]
		(*alarmDeal)[key]++
	}
}

func ReadXlsxFile(filename string) error {
	xlsxFile, err := excelize.OpenFile(filename)
	if err != nil {
		//return fmt.Errorf("read file:%s error:%v\n", filename, err)
		return err
	}
	rowAlarmList, _ := xlsxFile.GetRows("alarm_list")

	if rowAlarmList == nil {
		fmt.Printf("get %s sheet alarm_list fail\n", filename)
	} else {
		AnalysisAlarm(rowAlarmList, &rawAmarm)
	}

	rowAlarmDetail, _ := xlsxFile.GetRows("alarm_deal_list")
	if rowAlarmDetail == nil {
		fmt.Printf("get %s sheet alarm_deal_list fail\n", filename)
	} else {
		analysisDeal(rowAlarmDetail, &alarmDeal)
	}
	return nil
}

func main() {
	ReadXlsxFile("./alarm_data_20190417_20190425.xlsx")
	for k, v := range rawAmarm {
		fmt.Printf("%-100s\t%d\n", k, v)
	}
	for k, v := range alarmDeal {
		fmt.Printf("%-100s\t%d\n", k, v)
	}
	fmt.Println(rawAmarm)
	fmt.Println(alarmDeal)

}
