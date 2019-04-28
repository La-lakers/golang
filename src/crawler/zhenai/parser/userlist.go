package parser

import (
	"crawler/engine"
	"crawler/models"
	"log"
	"regexp"
)

const userListRe = `<div class="m-btn purple" data-v-bff6f798>([^<]+)</div>`

func UserProfile(content []byte, name, sex string) engine.ParseResult {
	re := regexp.MustCompile(userListRe)
	matchs := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	user := models.UserProfile{}
	if len(matchs) >= 9 {
		user.Name = name
		user.Sex = sex
		user.Marriage = exceptString(matchs[0])
		user.Age = exceptString(matchs[1])
		user.Xinzuo = exceptString(matchs[2])
		user.Height = exceptString(matchs[3])
		user.Weight = exceptString(matchs[4])
		user.Workplace = exceptString(matchs[5])
		user.Income = exceptString(matchs[6])
		user.Job = exceptString(matchs[7])
		user.Edu = exceptString(matchs[8])
		result.Items = append(result.Items, user)
	} else {
		log.Println()
	}

	return result
}

func exceptString(b [][]byte) string {
	if len(b) >= 2 {
		return string(b[1])
	}
	return "-"
}
