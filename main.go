package main

import (
	"esbuilder/esbuilder"
	"fmt"
)

func main() {
	boolQuery := esbuilder.NewBoolQuery()
	boolQuery.Filter(esbuilder.NewTermQuery("pnc_point.keyword", "liupeng"), esbuilder.NewRangeQuery("start_time").Gte(134123456789).Lte(134123456799))
	fmt.Println(boolQuery.BuildJson())
}
