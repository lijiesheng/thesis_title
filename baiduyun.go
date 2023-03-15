package main

import (
	"fmt"
	bdyp "github.com/zcxey2911/bdyp_upload_golang"
	"os"
)

/**
 * @Description
 * @Author lijiesheng
 * @Date 2023/3/13 12:49 PM
 **/

/**
* Appid、Appkey、Secretkey、Signkey是您应用实际开发的主要凭证，每个应用唯一标示，互不相同，请妥善保管。
  AppID：31227963
  Appkey：wvZNZMVpsyCeI4QzULG3KbGTyRIpe1qQ
  Secretkey：GN7F2O6wLGo7v63E0TcaLLzDfZUQa7MK
  Signkey：kC4-$$@Dvb338Nj!$YJZEf#f6FC=20Gs
  我的授权码: f13cd7f64209b3026e77a81d32ef1026   一次有效并在 10 分钟之后过期。
*/

// 百度网盘API 文档 https://pan.baidu.com/union/doc/fl0hhnulu

// 百度网盘API接入授权
// https://developer.aliyun.com/article/1168762?accounttraceid=9247f7cd31e64ba29cb73aeb75166a23ilax

/****  1、接入授权  *******/
// 注意授权码一次性有效并且会在10分钟后过期，随后编写代码获取token:
//授权码模式实现授权，主要依赖于以下 2 步：
//
//发起授权码 Code 请求，获取用户授权码 Code
//换取 Access Token 凭证

// 1. 通过浏览器获取 code
// 2. 获取 Access Token ,有效期是30天
// 3. 如果 Access Token 过期后怎么办？我们支持刷新 Access Token。那么，您如何刷新 Access Token呢。
func main() {

	var bcloud = bdyp.Bcloud{}

	// 获取token
	res, err := bcloud.GetToken("683375bb0c60dc1aa9ae13601520de28", "oob",
		"wvZNZMVpsyCeI4QzULG3KbGTyRIpe1qQ", "GN7F2O6wLGo7v63E0TcaLLzDfZUQa7MK")

	fmt.Println(res)

	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Printf("接口的token是: %#v\n", res.AccessToken)
	}
	// 读取文件
	f, err := os.Open("/Users/liuyue/Downloads/ju1.webp")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer f.Close()

	// 上传文件
	print(bcloud.Upload(&bdyp.FileUploadReq{
		Name:  "/apps/云盘备份/ju2.webp",
		File:  f,
		RType: nil,
	}))
}

// 文件下载
func