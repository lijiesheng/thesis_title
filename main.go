package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"thesis_title/downloadsdk/download"
	"thesis_title/model"
	"thesis_title/mysql"
	openapi "thesis_title/openxpanapi"
)

const (
	MyAccessToken = "121.1cfecd99c9e82abd64c23f26ccb1b719.YCwlE_VEjvelgcQe2wSrbTIHWDQxA5etpC87Tbp.3AL_qA"
)

/**
 * @Description
 * @Author lijiesheng
 * @Date 2023/1/10 4:52 PM
 * 框架 Gin
 **/
type DownRes struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {

	// 创建一个默认的路由引擎
	r := gin.New()
	r.Use(gin.Recovery())
	f, _ := os.OpenFile("./app01.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

	log.SetOutput(f)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
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
		count := 0

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

		if len(theisNameList) >= 4 {
			sql := `select id,title,size,type, size_int from thesis_title where ( title like "%.pdf" or title like "%.docx" or title like "%.doc") and title like ? and title like ? and title like ?  and title like ? limit ?,?`
			err = mysql.Db.Select(&thesisList, sql, "%"+theisNameList[0]+"%", "%"+theisNameList[1]+"%", "%"+theisNameList[2]+"%", "%"+theisNameList[3]+"%",
				(pageIndexInt-1)*pageNumberInt, pageNumberInt)
			if err != nil {
				fmt.Printf("query failed, err:%v\n", err)
				return
			}
			// 4、查询总数
			sql = `select count(*) count from thesis_title where ( title like "%.pdf" or title like "%.docx" or title like "%.doc") and title like ? and title like ? and title like ?  and title like ? `
			err = mysql.Db.Get(&count, sql, "%"+theisNameList[0]+"%", "%"+theisNameList[1]+"%", "%"+theisNameList[2]+"%", "%"+theisNameList[3]+"%")
			if err != nil {
				fmt.Printf("get failed, err:%v\n", err)
				return
			}
		}
		if len(theisNameList) == 3 {
			sql := `select id,title,size,type, size_int from thesis_title where ( title like "%.pdf" or title like "%.docx" or title like "%.doc") and title like ? and title like ? and title like ? limit ?,?`
			err = mysql.Db.Select(&thesisList, sql, "%"+theisNameList[0]+"%", "%"+theisNameList[1]+"%", "%"+theisNameList[2]+"%", (pageIndexInt-1)*pageNumberInt, pageNumberInt)
			if err != nil {
				fmt.Printf("query failed, err:%v\n", err)
				return
			}
			// 4、查询总数
			sql = `select count(*) count from thesis_title where ( title like "%.pdf" or title like "%.docx" or title like "%.doc") and title like ? and title like ? and title like ?`
			err = mysql.Db.Get(&count, sql, "%"+theisNameList[0]+"%", "%"+theisNameList[1]+"%", "%"+theisNameList[2]+"%")
			if err != nil {
				fmt.Printf("get failed, err:%v\n", err)
				return
			}
		}
		if len(theisNameList) == 2 {
			sql := `select id,title,size,type, size_int from thesis_title where ( title like "%.pdf" or title like "%.docx" or title like "%.doc") and title like ? and title like ? limit ?,?`
			err = mysql.Db.Select(&thesisList, sql, "%"+theisNameList[0]+"%", "%"+theisNameList[1]+"%", (pageIndexInt-1)*pageNumberInt, pageNumberInt)
			if err != nil {
				fmt.Printf("query failed, err:%v\n", err)
				return
			}
			// 4、查询总数
			sql = `select count(*) count from thesis_title where ( title like "%.pdf" or title like "%.docx" or title like "%.doc") and title like ? and title like ? `
			err = mysql.Db.Get(&count, sql, "%"+theisNameList[0]+"%", "%"+theisNameList[1]+"%")
			if err != nil {
				fmt.Printf("get failed, err:%v\n", err)
				return
			}
		}
		if len(theisNameList) == 1 {
			sql := `select id,title,size,type, size_int from thesis_title where ( title like '%.pdf' or title like '%.docx' or title like '%.doc') and title like ? limit ?,?`
			err = mysql.Db.Select(&thesisList, sql, "%"+theisNameList[0]+"%", (pageIndexInt-1)*pageNumberInt, pageNumberInt)
			if err != nil {
				fmt.Printf("query failed, err:%v\n", err)
				return
			}
			// 4、查询总数
			sql = `select count(*) count from thesis_title where ( title like "%.pdf" or title like "%.docx" or title like "%.doc") and title like ?`
			err = mysql.Db.Get(&count, sql, "%"+theisNameList[0]+"%")
			if err != nil {
				fmt.Printf("get failed, err:%v\n", err)
				return
			}
		}
		if len(theisNameList) == 0 {
			sql := `select id,title,size,type, size_int from thesis_title where ( title like "%.pdf" or title like "%.docx" or title like "%.doc") limit ?,?`
			err = mysql.Db.Select(&thesisList, sql, (pageIndexInt-1)*pageNumberInt, pageNumberInt)
			if err != nil {
				fmt.Printf("query failed, err:%v\n", err)
				return
			}

			// 4、查询总数
			sql = `select count(*) count from thesis_title where ( title like "%.pdf" or title like "%.docx" or title like "%.doc")`
			err = mysql.Db.Get(&count, sql)
			if err != nil {
				fmt.Printf("get failed, err:%v\n", err)
				return
			}
		}

		for i := 0; i < len(thesisList); i++ {
			titleSplitList := strings.Split(thesisList[i].Title, "/")
			thesisList[i].Title = titleSplitList[len(titleSplitList)-1]
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

	r.GET("/get_url", Cors, func(c *gin.Context) {
		// 1、获取文件名
		value := c.Query("search")
		log.Printf("downloadName=%s", value)
		if value == "" {
			return
		}
		// 2、在网盘中搜索文件
		xpanfilesearch := LjsMyXpanfilesearch(value)
		if len(xpanfilesearch.List) <= 0 {
			fmt.Println("没有需要下载的文件")
			return
		}
		fsids := "["
		for i := 0; i < len(xpanfilesearch.List) && i < 100; i++ {
			fsids += strconv.FormatUint(xpanfilesearch.List[i].Fsid, 10) + ","
		}
		fsids = fsids[0 : len(fsids)-1]
		fsids += "]"
		// 3、提供文件名和下载链接
		res := LjsXpanmultimediafilemetas(fsids)
		c.JSON(200, gin.H{
			"data":   res,
			"status": 200,
			"msg":    "获取成功",
		})
	})

	// 获取文件
	r.GET("/get_file", Cors, func(c *gin.Context) {
		c.Status(http.StatusNotFound)
		//// 获取当前文件路径
		//dir, _ := os.Getwd()
		//if path := c.Query("path"); path != "" {
		//	target := filepath.Join(dir, path)
		//	//target := filepath.Join(dir, "app01.log")
		//	log.Println(target)
		//	c.Header("Content-Description", "File Transfer")
		//	c.Header("Content-Transfer-Encoding", "binary")
		//	c.Header("Content-Disposition", "attachment;filename="+path)
		//	c.Header("content-type", "application/octet-stream;charset=utf-8")
		//	c.File(target)
		//} else {
		//	c.Status(http.StatusNotFound)
		//}
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

// 搜索文件 string |  搜索的关键字
// 获取文件列表
func LjsMyXpanfilesearch(key string) download.FileMetasReturn {
	// 1、搜索文件
	accessToken := MyAccessToken // string |
	web := "1"                   // string |  (optional)
	num := "500"                 // string |  默认为500，不能修改
	page := "1"                  // string |  页数，从1开始，缺省则返回所有条目
	dir := "/"                   // string |  搜索目录，默认根目录
	recursion := "1"             // string |  是否递归搜索子目录 1:是，0:否（默认）
	configuration := openapi.NewConfiguration()
	api_client := openapi.NewAPIClient(configuration)
	_, r, err := api_client.FileinfoApi.Xpanfilesearch(context.Background()).
		AccessToken(accessToken).Web(web).Num(num).Page(page).
		Dir(dir).Recursion(recursion).Key(key).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileinfoApi.Xpanfilesearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", r)
	}

	fileMetasData := download.FileMetasReturn{}
	err = json.Unmarshal(bodyBytes, &fileMetasData)

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", r)
	}
	return fileMetasData

	//bodyString := string(bodyBytes)
	//fmt.Fprintf(os.Stdout, "Response from `FileinfoApi.Xpanfilesearch`: %v body: %v\n", r, bodyString)
	//fmt.Fprintf(os.Stdout, "Response from `FileinfoApi.Xpanfilesearch`: %v\n", resp)
}

// 查询文件信息 获取下载地址dlink
func LjsXpanmultimediafilemetas(fsidstr string) []DownRes {
	accessToken := MyAccessToken // string
	thumb := "1"                 // string |  (optional)
	extra := "1"                 // string |  (optional)
	//fsids := "[1103236387625589,1101724761997348]" // 文件id数组，数组中元素是uint64类型，数组大小上限是：100
	dlink := "1" // 是否需要下载地址，0为否，1为是，默认为0。获取到dlink后，参考下载文档进行下载操作
	//aa := strings(fsids)

	configuration := openapi.NewConfiguration()
	api_client := openapi.NewAPIClient(configuration)
	_, r, err := api_client.MultimediafileApi.Xpanmultimediafilemetas(context.Background()).AccessToken(accessToken).
		Thumb(thumb).Extra(extra).Fsids(fsidstr).Dlink(dlink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultimediafileApi.Xpanmultimediafilemetas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	fileMetasData := download.FileMetasReturn{}
	err = json.Unmarshal(bodyBytes, &fileMetasData)

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", r)
	}
	d := []DownRes{}

	for _, v := range fileMetasData.List {
		// 下载文件
		//download.Download(accessToken, v.Dlink, v.Filename)
		download.Download(accessToken, v.Dlink, "../logs/fanlai/"+v.Filename)
		// 获取地址
		//pwd, _ := os.Getwd()
		// 拼接数据
		d = append(d, DownRes{
			Name: v.Filename,
			Url:  v.Filename,
		})
	}
	return d
}
