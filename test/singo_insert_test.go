package test

import (
	"colume-search/db"
	"colume-search/inputData"
	"colume-search/utils"
	"fmt"
	"testing"
)

func TestSingoInsert(t *testing.T) {
	dirPath := "/Users/hejinlong/Desktop/tda/demo/export/"
	files := utils.Walk(dirPath)
	table := db.GetTable("test")
	for _, filePath := range files {
		fmt.Println(filePath)
		inData := &inputData.InputData{}
		utils.LoadJson(filePath, inData)
		table.Insert(inData)
	}

	table.Save()
}