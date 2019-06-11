// config.go
package utils

import (
	ini "gopkg.in/ini.v1"
)

var fileReader *ini.File

func init() {
	opt := ini.LoadOptions{AllowBooleanKeys: true}
	cfg, err := ini.LoadSources(opt, "config.ini")
	if err != nil {
		panic(err)
	}
	fileReader = cfg
}

func GetKeyBySlice(section string) []string {
	s, err := fileReader.GetSection(section)
	if err != nil {
		panic(err)
	}
	return s.KeyStrings()
}
