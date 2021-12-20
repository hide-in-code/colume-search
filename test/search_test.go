package test

import (
	"colume-search/db"
	"colume-search/utils"
	"fmt"
	"testing"
)

func TestSampleSearch(t *testing.T) {
	table := db.GetTable("test")
	res, err := table.Search("index", "1")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func TestMultiSearch(t *testing.T)  {
	table := db.GetTable("test")
	res, err := table.MultiSearch("index", []string{"1", "2"})
	if err != nil {
		panic(err)
	}

	fmt.Println(utils.ArrayUnique(res))
}