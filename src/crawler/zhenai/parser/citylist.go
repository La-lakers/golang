package parser

import (
	"crawler/engine"
	"regexp"
)

const citylistRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z]+)[^>]*>([^<]+)`

//const citylistRe = `<a href="(http://www.zhenai.com/zhenghun/qingyang2)[^>]*>([^<]+)`

func ParseCityList(content []byte) engine.ParseResult {
	re := regexp.MustCompile(citylistRe)
	matchs := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	limit := 0
	for _, m := range matchs {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:       string(m[1]),
				ParseFunc: ParseCity,
			})
		result.Items = append(result.Items, string(m[2]))
		limit++
		if limit > 5 {
			break
		}
	}
	return result

}
