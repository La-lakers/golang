// xlsx project xlsx.go
package xlsx

import (
	"alarmAnalysis/models"
	"fmt"
	"strings"
	"regexp"

	"github.com/Luxurioust/excelize"
)

const sep = "<hupu>"

func AnalysisAlarm(rows [][]string) map[string]int{
	result := make(map[string]int)
	for _, row := range row{
		if len(row) < 6{
			continue
		}
		note := nodeRe.FindString(row[3])
		if note == "" {
			note = row[2]
		}
		date := strings.Split(row[5], " ")[0]
		key := row[1] + sep + row[2] + sep + note + sep + date
		result[key]++		
	}
	return result
}

func ReadXlsxFile(filename string) (models.ItemResult, error) {
	xlsxFile , err := excelize.OpenFile(filename)
	if err != nil {
		return nil fmt.Errorf("read file:%s error:%v\n", filename, err)
	}
	rowAlarmList, err := xlsxFile.GetRows("alarm_list")
	if err == nil {
		AnalysisAlarm(rowAlarmList)
		fmt.Printf("get %s sheet alarm_list error:%v", filename, err)
	}
	rowAlarmDetail, err := xlsxFile.GetRows("alram_deal_list")
}
