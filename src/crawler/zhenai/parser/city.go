// city
package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/\d+)"[^>]*>([^<]+).*?性别：</span>(女士)`

func ParseCity(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matchs := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}

	for _, m := range matchs {
		name := string(m[2])
		sex := string(m[3])
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParseFunc: func(c []byte) engine.ParseResult {
					return UserProfile(c, name, sex)
				}})
		result.Items = append(result.Items, string(m[2]))

	}
	return result

}
