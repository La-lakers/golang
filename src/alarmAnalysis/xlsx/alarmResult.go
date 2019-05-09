// alarmResult.go
package xlsx

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Luxurioust/excelize"
)

var (
	sep     = "<hupu>"
	noteRe  = regexp.MustCompile(`note\[[^]]+]`)
	trouble = []string{"有影响",
		"影响业务",
		"有问题",
		"有异常",
		"有质量影响",
		"影响卡顿"}
)

type AlarmList struct {
	Service string
	Item    string
	Content string
	Count   int
	Date    string
}

type AlarmDealList struct {
	Service string
	Item    string
	Handle  string
	Result  string
	Count   int
}

type AlarmResult struct {
	AlarmLists     []AlarmList
	AlarmDealLists []AlarmDealList
	FileName       string
	Xlsx           *excelize.File
}

func NewAlarmResult(filename string) (AlarmResult, error) {
	var alarmResult AlarmResult
	alarmResult.FileName = filename
	xlsx, err := excelize.OpenFile(alarmResult.FileName)
	if err != nil {
		fmt.Println(filename, err)
		return alarmResult, err
	}
	alarmResult.Xlsx = xlsx
	alarmResult.SetAlarmList()
	alarmResult.SetAlarmDeal()
	return alarmResult, nil
}

func (p *AlarmResult) SetAlarmList() {
	alarmList, err := p.Xlsx.GetRows("alarm_list")
	if err != nil || alarmList == nil {
		fmt.Printf("get sheet alarm_list of %s err:%v", p.FileName, err)
		return
	}
	result := make(map[string]int)
	alarm := AlarmList{}
	var alarmLists []AlarmList
	for _, row := range alarmList[1:] {
		if len(row) < 6 {
			continue
		}
		note := noteRe.FindString(row[3])
		if note == "" {
			note = row[2]
		}
		date := strings.Split(row[5], " ")[0]
		key := row[1] + sep + row[2] + sep + note + sep + date
		result[key]++
	}
	for k, v := range result {
		key := strings.Split(k, sep)
		alarm.Service = key[0]
		alarm.Item = key[1]
		alarm.Content = key[2]
		alarm.Count = v
		alarm.Date = key[3]
		alarmLists = append(alarmLists, alarm)
	}
	p.AlarmLists = alarmLists
}

func (p *AlarmResult) SetAlarmDeal() {
	alarmDeal, err := p.Xlsx.GetRows("alarm_deal_list")
	if err != nil || alarmDeal == nil {
		fmt.Printf("get sheet alarm_deal_list of %s  is empty or error:%v", p.FileName, err)
		return
	}
	result := make(map[string]int)
	alarm := AlarmDealList{}
	var alarmDealLists []AlarmDealList
	for _, row := range alarmDeal[1:] {
		if len(row) < 4 {
			continue
		}
		key := row[1] + sep + row[2] + sep + row[3]
		result[key]++
	}
	for k, v := range result {
		effect := "无影响"
		key := strings.Split(k, sep)
		alarm.Service = key[0]
		alarm.Item = key[1]
		alarm.Handle = key[2]
		alarm.Count = v
		for _, s := range trouble {
			if strings.Contains(key[2], s) {
				effect = "有影响"
				break
			}
		}
		alarm.Result = effect
		alarmDealLists = append(alarmDealLists, alarm)
	}
	p.AlarmDealLists = alarmDealLists
}
