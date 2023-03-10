package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
	"os"
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
	r := gin.New()
	r.Use(gin.Recovery())
	f, _ := os.OpenFile("./app01.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

	var conf = gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("客户端IP:%s,请求时间:[%s],请求方式:%s,请求地址:%s,http协议版本:%s,请求状态码:%d,响应时间:%s,客户端:%s，错误信息:%s\n",
				param.ClientIP,
				param.TimeStamp.Format("2006年01月02日 15:03:04"),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		},
		Output: io.MultiWriter(os.Stdout, f),
	}
	r.Use(gin.LoggerWithConfig(conf))

	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.GET("/get_theis_title", Cors, func(c *gin.Context) {
		// 1、获取参数
		theisName := c.DefaultQuery("theis_name", "单片机")
		log.Printf("theisName=%s", theisName)
		pageNumber := c.DefaultQuery("page_number", "20")
		log.Printf(",pageNumber=%s", pageNumber)
		pageIndex := c.DefaultQuery("page_index", "1")
		log.Printf(",pageIndex=%s", pageIndex)

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
		theisNameList := strings.Fields(theisName)
		sql := ``

		if len(theisNameList) >= 4 {
			sql := `select id,title,size,type, size_int from thesis_title where title like ? and title like ? and title like ?  and title like ? limit ?,?`
			err = mysql.Db.Select(&thesisList, sql, "%"+theisNameList[0]+"%", "%"+theisNameList[1]+"%", "%"+theisNameList[2]+"%", "%"+theisNameList[3]+"%",
				(pageIndexInt-1)*pageNumberInt, pageNumberInt)
			if err != nil {
				fmt.Printf("query failed, err:%v\n", err)
				return
			}
		}
		if len(theisNameList) == 3 {
			sql := `select id,title,size,type, size_int from thesis_title where title like ? and title like ? and title like ? limit ?,?`
			err = mysql.Db.Select(&thesisList, sql, "%"+theisNameList[0]+"%", "%"+theisNameList[1]+"%", "%"+theisNameList[2]+"%", (pageIndexInt-1)*pageNumberInt, pageNumberInt)
			if err != nil {
				fmt.Printf("query failed, err:%v\n", err)
				return
			}
		}
		if len(theisNameList) == 2 {
			sql := `select id,title,size,type, size_int from thesis_title where title like ? and title like ? limit ?,?`
			err = mysql.Db.Select(&thesisList, sql, "%"+theisNameList[0]+"%", "%"+theisNameList[1]+"%", (pageIndexInt-1)*pageNumberInt, pageNumberInt)
			if err != nil {
				fmt.Printf("query failed, err:%v\n", err)
				return
			}
		}
		if len(theisNameList) == 1 {
			sql := `select id,title,size,type, size_int from thesis_title where title like ? limit ?,?`
			err = mysql.Db.Select(&thesisList, sql, "%"+theisNameList[0]+"%", (pageIndexInt-1)*pageNumberInt, pageNumberInt)
			if err != nil {
				fmt.Printf("query failed, err:%v\n", err)
				return
			}
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
			"pageNumber": pageNumberInt,
			"pageIndex":  pageIndexInt,
			"status":     200,
			"msg":        "获取成功",
		})
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}

func Cors(context *gin.Context) {
	method := context.Request.Method
	// 必须，接受指定域的请求，可以使用*不加以限制，但不安全
	//context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Origin", context.GetHeader("Origin"))
	fmt.Println(context.GetHeader("Origin"))
	// 必须，设置服务器支持的所有跨域请求的方法
	context.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
	context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
	// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
	context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
	context.Header("Access-Control-Allow-Credentials", "true")
	// 放行所有OPTIONS方法
	if method == "OPTIONS" {
		context.AbortWithStatus(http.StatusNoContent)
		return
	}
	context.Next()
}
