// models project models.go
package models

type AlarmList struct {
	Service string
	Item    string
	Content string
	Count   int
	Date    string
}

type AlarmDealList struct {
	Service    string
	Item       string
	Trouble    int
	Count      int
	Percentage float32
}

type ItemResult struct {
	RawAlarms     []RawAlarm
	HandleDetails []HandleDetail
}
