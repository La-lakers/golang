package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)

func main() {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}
	exists, err := client.IndexExists("index").Do(ctx)
	if err != nil {
		panic(err)
	}
	if exists {
		fmt.Println("ok")
	}
	esquery := elastic.NewBoolQuery()
	esquery.Must(elastic.NewQueryStringQuery("*"))
	query := elastic.NewRangeQuery("createDate")
	query.Gte(1559809963191)
	query.Lte(1559824363191)
	esquery.Filter(query)
	fmt.Println(esquery.Source())

	get1, err := client.Search().Index("index").Query(esquery).From(0).Size(10).Pretty(true).Do(ctx)
	if err != nil {
		fmt.Println("error")
		panic(err)
	}
	fmt.Println(get1.TookInMillis)

}
