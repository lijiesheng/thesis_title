package main

import (
	"fmt"
	"thesis_title/downloadsdk/download"
	//"icode.baidu.com/baidu/xpan/go-sdk/xpan/download"
)

func main() {
	// 使用示例

	// 用户的access_token
	accessToken := "your-access-token"
	accessToken = "123.56c5d1f8eedf1f9404c547282c5dbcf4.YmmjpAlsjUFbPly3mJizVYqdfGDLsBaY5pyg3qL.a9IIIQ"

	// 要查询的文件的fsid，支持多个文件查询
	var fsids []uint64
	// fsids = append(fsids, 954192397535589) //文档fsid
	// fsids = append(fsids, 374004028691511) // 图片fsid
	// fsids = append(fsids, 17839481439482)  //音频fisd
	fsids = append(fsids, 89922956254800) //视频fisd

	// 一般不需要，这里使用时填空
	// 如果是查询共享目录或专属空间内文件时需要path，可结合官网文档
	path := ""

	// call Api
	arg := download.NewFileMetasArg(fsids, path)
	ret, err := download.FileMetas(accessToken, arg)
	if err != nil {
		fmt.Printf("[msg: filemetas error] [err:%v]", err.Error())
	} else {
		fmt.Printf("ret:%+v", ret)
		fmt.Printf("ret.List:%+v", ret.List)
		// 获取list的第一个元素的dlink示例
		fmt.Printf(ret.List[0].Dlink)
	}
}
