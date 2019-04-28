// models project models.go
package models

type RawAlarm struct {
	Service string
	Item    string
	Content string
	Count   int
	Date    string
}

type HandleDetail struct {
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
