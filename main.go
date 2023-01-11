package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"thesis_title/model"
	"thesis_title/mysql"
)

/**
 * @Description
 * @Author lijiesheng
 * @Date 2023/1/10 4:52 PM
 * 框架 Gin
 **/
func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.GET("/get_theis_title", func(c *gin.Context) {
		// 1、获取参数
		theisName := c.DefaultQuery("theis_name", "单片机")
		pageNumber := c.DefaultQuery("page_number", "20")
		pageIndex := c.DefaultQuery("page_index", "1")

		// 2、整理参数
		pageIndexInt, err := strconv.Atoi(pageIndex)
		if err != nil {
			fmt.Println(err)
			return
		}
		pageNumberInt, err := strconv.Atoi(pageNumber)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 3、返回结果
		var thesisList []model.Thesis
		sql := `select id,title,size,type, size_int from thesis_title where title like ? limit ?,?`
		err = mysql.Db.Select(&thesisList, sql, "%"+theisName+"%", (pageIndexInt-1)*pageNumberInt, pageNumberInt)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return
		}
		for i := 0; i < len(thesisList); i++ {
			titleSplitList := strings.Split(thesisList[i].Title, "/")
			thesisList[i].Title = titleSplitList[len(titleSplitList)-1]
		}

		// 4、查询总数
		var count int
		sql = `select count(*) count from thesis_title where title like ?`
		err = mysql.Db.Get(&count, sql, "%"+theisName+"%")
		if err != nil {
			fmt.Printf("get failed, err:%v\n", err)
			return
		}

		c.JSON(200, gin.H{
			"data":       thesisList,
			"total":      count,
			"pageNumber": pageNumber,
			"pageIndex":  pageIndexInt,
		})
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
