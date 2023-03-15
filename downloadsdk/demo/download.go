package main

import (
	"fmt"
	"thesis_title/downloadsdk/download"
	//"icode.baidu.com/baidu/xpan/go-sdk/xpan/download"
)

func main() {

	// 使用示例

	// 用户的access_token
	accessToken := "123.56c5d1f8eedf1f9404c547282c5dbcf4.YmmjpAlsjUFbPly3mJizVYqdfGDLsBaY5pyg3qL.a9IIIQ"

	// *** 仅针对小文件，大文件可参考进行改造升级 ***
	// 下载地址 dlink 示例
	// 一定要保证 dlink 正确
	// dlink有效期为8小时，过期后，dlink失效
	// 通过查询文件信息接口（filemetas）获取到 dlink。
	// 查询文件信息接口（filemetas）通过获取文件列表接口拿到文件fsid。获取文件列表已附上go-sdk，见官网文档页参考使用
	dlink := "https://d.pcs.baidu.com/file/9c70b497bde60464d12911279473dce6?fid=1027554777-250528-374004028691511&rt=pr&sign=FDtAERV-DCb740ccc5511e5e8fedcff06b081203-PRqVwEvB%2BTSf%2BBlhFPoxsPm7QUo%3D&expires=8h&chkbd=0&chkv=2&dp-logid=942960704440686273&dp-callid=0&dstime=1652669893&r=173264414&origin_appid=25809617&file_type=0"

	// 下载数据输出到文件，可自定义文件名
	outputFilename := "output"

	// call API
	err := download.Download(accessToken, dlink, outputFilename)
	if err != nil {
		fmt.Printf("[msg: download error] [err:%v]", err.Error())
	}
}
