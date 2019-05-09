// alarmAnalysis project main.go
package main

import (
	"alarmAnalysis/xlsx"
	"flag"
	"fmt"
	"strings"
	"time"
)

func main() {
	var fileString string
	flag.StringVar(&fileString, "f", "", "xlsx filenames")
	flag.Parse()
	if fileString == "" {
		flag.Usage()
		return
	}
	filenames := strings.Split(fileString, ",")
	var alarmResultList xlsx.AlarmResultList

	for _, filename := range filenames {
		fmt.Println(filename)
		alarmResult, err := xlsx.NewAlarmResult(filename)
		if err != nil {
			fmt.Printf("open filename:%s error:%v\n", filename, err)
			continue
		}
		alarmResultList = append(alarmResultList, alarmResult)
	}
	fmt.Println(len(alarmResultList))
	saveFilename := fmt.Sprintf("alarm-%s.xlsx", time.Now().Format("2006-01-02"))
	(&alarmResultList).Save(saveFilename)
}
