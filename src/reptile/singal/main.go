// singal project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const url = "http://www.zhenai.com/zhenghun"

func getURLByCity(text string) {
	re := regexp.MustCompile(`<a href="(https?://.*?)".*?>(.*?)</a>`)
	result := re.FindAllStringSubmatch(text, -1)
	for _, i := range result {
		if len(i) >= 3 {
			fmt.Printf("%s\t%s\n", i[2], i[1])
		}
	}

}

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("%s response status code(%d) not 200",
			url, resp.StatusCode)
		return
	}
	allCityPage, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", allCityPage)
	getURLByCity(string(allCityPage))
}
