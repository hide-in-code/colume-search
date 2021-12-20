package controller

import (
	"colume-search/db"
	"colume-search/inputData"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Add(c *gin.Context) {
	tableName := c.Param("table")
	jsonData := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&jsonData)

	if _, ok := jsonData["key"];!ok {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  201,
				"data":    "",
				"message": "请求数据格式不对，样例：http://xxxx.xxx.xx/demo.json",
			},
		)
		return
	}

	if _, ok := jsonData["data"];!ok {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  201,
				"data":    "",
				"message": "请求数据格式不对，样例：http://xxxx.xxx.xx/demo.json",
			},
		)
		return
	}

	postKey, ok := jsonData["key"].(string)
	if !ok {//断言失败
		c.JSON(
			http.StatusOK, gin.H{
				"status":  201,
				"data":    "",
				"message": "请求数据格式不对，[key]的值必须是字符串类型，样例：http://xxxx.xxx.xx/demo.json",
			},
		)
		return
	}

	table := db.GetTable(tableName)
	inData := &inputData.InputData{
		Key: postKey,
		Data: jsonData["data"],
	}

	_, err := table.Insert(inData)
	if err != nil {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  203,
				"data":    []string{},
				"message": err,
			},
		)
		return
	}

	table.Save()

	c.JSON(
		http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    "",
			"message": "ok",
		},
	)
}

func Search(c *gin.Context)  {
	tableName := c.Param("table")
	indexName := c.DefaultQuery("index", "")
	valueName := c.DefaultQuery("value", "")
	isCount := c.DefaultQuery("isCount", "")

	if indexName == "" {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  201,
				"data":    []string{},
				"message": "必须为查询指定一个索引，用法：/:table?index=index1&value=value1",
			},
		)
		return
	}

	if valueName == "" {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  201,
				"data":    []string{},
				"message": "必须为查询指定一个值，用法：/:table?index=index1&value=value1",
			},
		)
		return
	}

	table := db.GetTable(tableName)
	if !table.CheckIndexExist(indexName) {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  202,
				"data":    []string{},
				"message": "index[" + indexName + "]不存在",
			},
		)
		return
	}

	res, err := table.Search(indexName, valueName)
	if err != nil {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  203,
				"data":    []string{},
				"message": err,
			},
		)
		return
	}

	if isCount != "" {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  203,
				"data":    len(res),
				"message": "",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  203,
			"data":    res,
			"message": "",
		},
	)
}

func AllIndex(c *gin.Context)  {
	tableName := c.Param("table")
	table := db.GetTable(tableName)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    table.AllIndex(50),
			"message": "ok",
		},
	)
}

func AllIndexCount(c *gin.Context)  {
	tableName := c.Param("table")
	table := db.GetTable(tableName)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    table.AllIndexCount(),
			"message": "ok",
		},
	)
}