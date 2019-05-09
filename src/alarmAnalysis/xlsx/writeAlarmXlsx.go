// writeAlarmXlsx.go
package xlsx

import (
	"fmt"
	"os"

	"github.com/Luxurioust/excelize"
)

type AlarmResultList []AlarmResult

func (p *AlarmResultList) wirteList(file *excelize.File) {
	lineNum := 2
	sheetName := "报警详情"
	file.SetCellValue(sheetName, "A1", "业务线")
	file.SetCellValue(sheetName, "B1", "报警项")
	file.SetCellValue(sheetName, "C1", "报警内容")
	file.SetCellValue(sheetName, "D1", "报警个数")
	file.SetCellValue(sheetName, "E1", "日期")
	for _, alarmResultList := range *p {
		for _, a := range alarmResultList.AlarmLists {
			A := fmt.Sprintf("A%d", lineNum)
			file.SetCellValue(sheetName, A, a.Service)
			B := fmt.Sprintf("B%d", lineNum)
			file.SetCellValue(sheetName, B, a.Item)
			C := fmt.Sprintf("C%d", lineNum)
			file.SetCellValue(sheetName, C, a.Content)
			D := fmt.Sprintf("D%d", lineNum)
			file.SetCellValue(sheetName, D, a.Count)
			E := fmt.Sprintf("E%d", lineNum)
			file.SetCellValue(sheetName, E, a.Date)
			lineNum++
		}
	}
}

func (p *AlarmResultList) wirteDeal(file *excelize.File) {
	lineNum := 2
	sheetName := "报警处理详情"
	file.SetCellValue(sheetName, "A1", "业务线")
	file.SetCellValue(sheetName, "B1", "报警项")
	file.SetCellValue(sheetName, "C1", "处理内容")
	file.SetCellValue(sheetName, "D1", "报警影响")
	file.SetCellValue(sheetName, "E1", "报警数量")
	for _, alarmResultList := range *p {
		for _, a := range alarmResultList.AlarmDealLists {
			A := fmt.Sprintf("A%d", lineNum)
			file.SetCellValue(sheetName, A, a.Service)
			B := fmt.Sprintf("B%d", lineNum)
			file.SetCellValue(sheetName, B, a.Item)
			C := fmt.Sprintf("C%d", lineNum)
			file.SetCellValue(sheetName, C, a.Handle)
			D := fmt.Sprintf("D%d", lineNum)
			file.SetCellValue(sheetName, D, a.Result)
			E := fmt.Sprintf("E%d", lineNum)
			file.SetCellValue(sheetName, E, a.Count)
			lineNum++
		}
	}
}

func (p *AlarmResultList) Save(filename string) {
	xlsx := excelize.NewFile()
	xlsx.SetSheetName("Sheet1", "报警详情")
	_ = xlsx.NewSheet("报警处理详情")
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("open %s error:%v\n", filename, err))
	}
	defer file.Close()
	p.wirteList(xlsx)
	p.wirteDeal(xlsx)
	xlsx.WriteTo(file)
}
