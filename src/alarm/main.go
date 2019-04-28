// alarm project main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/Luxurioust/excelize"
)

const (
	nodeList  = `note\[[^]]+]`
	separator = `|hupu|`
	//filename = "./alarm.xlsx"
)

type Alarm struct {
	Service string
	Item    string
	Note    string
	Count   int
	Date    string
}

var xlxsFilename = fmt.Sprintf("anaylse-%s.xlsx", time.Now().Format("2006-01-02"))

func createExcel(a []Alarm) {
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A1", "业务线")
	xlsx.SetCellValue("Sheet1", "B1", "报警项")
	xlsx.SetCellValue("Sheet1", "C1", "报警内容")
	xlsx.SetCellValue("Sheet1", "D1", "报警个数")
	xlsx.SetCellValue("Sheet1", "E1", "日期")
	for i := 0; i < len(a); i++ {
		A := fmt.Sprintf("A%d", i+2)
		xlsx.SetCellValue("Sheet1", A, a[i].Service)
		B := fmt.Sprintf("B%d", i+2)
		xlsx.SetCellValue("Sheet1", B, a[i].Item)
		C := fmt.Sprintf("C%d", i+2)
		xlsx.SetCellValue("Sheet1", C, a[i].Note)
		D := fmt.Sprintf("D%d", i+2)
		xlsx.SetCellValue("Sheet1", D, a[i].Count)
		E := fmt.Sprintf("E%d", i+2)
		xlsx.SetCellValue("Sheet1", E, a[i].Date)
	}
	file, err := os.OpenFile(xlxsFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	xlsx.WriteTo(file)

}

func ParseAlarmXlsx(filename string) ([]Alarm, error) {
	result := make(map[string]int)
	var alarms []Alarm
	alarm := Alarm{}
	xlsx, err := excelize.OpenFile(filename)
	if err != nil {
		//log.Printf("read filename:%s error:%v", filename, err)
		return nil, err
	}
	//sheets := xlsx.GetSheetMap()
	rows, err := xlsx.GetRows("alarm_list")
	if err != nil {
		//log.Printf("get row error:%v", err)
		return nil, err
	}
	nodeRe := regexp.MustCompile(nodeList)
	for _, row := range rows[1:] {
		if len(row) < 6 {
			continue
		}
		note := nodeRe.FindString(row[3])
		if note == "" {
			note = row[2]
		}
		date := strings.Split(row[5], " ")[0]
		key := row[1] + separator + row[2] + separator + note + separator + date
		result[key]++
	}
	for k, v := range result {
		key := strings.Split(k, separator)
		alarm.Service = key[0]
		alarm.Item = key[1]
		alarm.Note = key[2]
		alarm.Count = v
		alarm.Date = key[3]
		alarms = append(alarms, alarm)
	}
	return alarms, nil

}

func main() {
	var fileString string
	flag.StringVar(&fileString, "f", "", "xlsx filenames")
	flag.Parse()
	if fileString == "" {
		flag.Usage()
	}
	filenames := strings.Split(fileString, ",")
	//alarm := Alarm{}
	var resultAlarm []Alarm
	for _, filename := range filenames {
		result, err := ParseAlarmXlsx(filename)
		if err != nil {
			log.Printf("Parse filename:%s error:%v", filename, err)
			continue
		}
		resultAlarm = append(resultAlarm, result...)
	}
	createExcel(resultAlarm)
}
