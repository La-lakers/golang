// engine project engine.go
package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var q []Request
	for _, re := range seeds {
		q = append(q, re)
	}
	for len(q) > 0 {
		request := q[0]
		q = q[1:]
		log.Printf("fetching url: %s", request.Url)
		body, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Printf("get url:%s, error:%v", request.Url, err)
			continue
		}
		data := request.ParseFunc(body)
		q = append(q, data.Requests...)
		for _, item := range data.Items {
			log.Printf("Got item %s", item)
		}
	}

}
