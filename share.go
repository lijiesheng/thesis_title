//package main
//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"os"
//	"strconv"
//	"thesis_title/downloadsdk/download"
//	openapi "thesis_title/openxpanapi"
//)
//
///**
// * @Description
// * @Author lijiesheng
// * @Date 2023/3/14 3:48 PM
// * 文件分享
// **/
//
//const (
//	MyAccessToken = "121.1cfecd99c9e82abd64c23f26ccb1b719.YCwlE_VEjvelgcQe2wSrbTIHWDQxA5etpC87Tbp.3AL_qA"
//)
//
//// 先做下载
//func main() {
//	// 1、搜索文件 获取文件的fsid
//	xpanfilesearchList := LjsMyXpanfilesearch("BMS")
//	if len(xpanfilesearchList.List) <= 0 {
//		fmt.Println("没有需要下载的文件")
//		return
//	}
//	fsids := "["
//	for i := 0; i < len(xpanfilesearchList.List) && i < 100; i++ {
//		fsids += strconv.FormatUint(xpanfilesearchList.List[i].Fsid, 10) + ","
//	}
//	fsids = fsids[0 : len(fsids)-1]
//	fsids += "]"
//	// 2、查询文件信息，获取对应的下载地址
//	LjsXpanmultimediafilemetas(fsids)
//}
//
//// 搜索文件 string |  搜索的关键字
//// 获取文件列表
//func LjsMyXpanfilesearch(key string) download.FileMetasReturn {
//	// 1、搜索文件
//	accessToken := MyAccessToken // string |
//	web := "1"                   // string |  (optional)
//	num := "500"                 // string |  默认为500，不能修改
//	page := "1"                  // string |  页数，从1开始，缺省则返回所有条目
//	dir := "/"                   // string |  搜索目录，默认根目录
//	recursion := "1"             // string |  是否递归搜索子目录 1:是，0:否（默认）
//	configuration := openapi.NewConfiguration()
//	api_client := openapi.NewAPIClient(configuration)
//	_, r, err := api_client.FileinfoApi.Xpanfilesearch(context.Background()).
//		AccessToken(accessToken).Web(web).Num(num).Page(page).
//		Dir(dir).Recursion(recursion).Key(key).Execute()
//
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Error when calling `FileinfoApi.Xpanfilesearch``: %v\n", err)
//		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
//	}
//
//	bodyBytes, err := ioutil.ReadAll(r.Body)
//
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "err: %v\n", r)
//	}
//
//	fileMetasData := download.FileMetasReturn{}
//	err = json.Unmarshal(bodyBytes, &fileMetasData)
//
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "err: %v\n", r)
//	}
//	return fileMetasData
//
//	//bodyString := string(bodyBytes)
//	//fmt.Fprintf(os.Stdout, "Response from `FileinfoApi.Xpanfilesearch`: %v body: %v\n", r, bodyString)
//	//fmt.Fprintf(os.Stdout, "Response from `FileinfoApi.Xpanfilesearch`: %v\n", resp)
//}
//
//// 查询文件信息 获取下载地址dlink
//func LjsXpanmultimediafilemetas(fsidstr string) download.FileMetasReturn {
//	accessToken := MyAccessToken // string
//	thumb := "1"                 // string |  (optional)
//	extra := "1"                 // string |  (optional)
//	//fsids := "[1103236387625589,1101724761997348]" // 文件id数组，数组中元素是uint64类型，数组大小上限是：100
//	dlink := "1" // 是否需要下载地址，0为否，1为是，默认为0。获取到dlink后，参考下载文档进行下载操作
//	//aa := strings(fsids)
//
//	configuration := openapi.NewConfiguration()
//	api_client := openapi.NewAPIClient(configuration)
//	_, r, err := api_client.MultimediafileApi.Xpanmultimediafilemetas(context.Background()).AccessToken(accessToken).
//		Thumb(thumb).Extra(extra).Fsids(fsidstr).Dlink(dlink).Execute()
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Error when calling `MultimediafileApi.Xpanmultimediafilemetas``: %v\n", err)
//		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
//	}
//	bodyBytes, err := ioutil.ReadAll(r.Body)
//	fileMetasData := download.FileMetasReturn{}
//	err = json.Unmarshal(bodyBytes, &fileMetasData)
//
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "err: %v\n", r)
//	}
//	return fileMetasData
//}
