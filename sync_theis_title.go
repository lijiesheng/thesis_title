package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"strings"
	"thesis_title/mysql"
	"time"
)

/**
 * @Description
 * @Author lijiesheng
 * @Date 2023/1/8 8:49 PM
 * 论文名称
 * https://blog.csdn.net/HYZX_9987/article/details/100072442
 **/

func main() {
	//GetFiles("/Users/mac/Downloads/论文/成稿1库/21年1万份稿子")
	//GetFiles("/Users/mac/Downloads/论文/2021年成稿库2.zip等多个文件")
}

// 获取文件名称
func GetFiles(folder string) {
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			GetFiles(folder + "/" + file.Name())
		} else {
			//fmt.Print(file.Name())
			//fmt.Print("  ")
			//size, s := getFileSize(file.Size())
			//fmt.Print(size)
			//fmt.Print(s)
			//fmt.Print("  ")
			//fmt.Print(file.ModTime().Format("2006-01-02 15:04:05"))
			//fmt.Print("  ")
			//fmt.Print(getFileType(file.Name()))
			//
			//fmt.Println()
			size, s := getFileSize(file.Size())

			// 写入数据库
			insertSql := `insert into thesis_title(title, size, type,year, size_int, create_time) values(?, ?, ?, ?, ?, ?)`
			_, err := mysql.Db.Exec(insertSql, folder+"/"+file.Name(), size, s, 2022, file.Size(), time.Now())
			if err != nil {
				fmt.Printf("get lastinsert ID failed, err:%v\n", err)
				return
			}
		}
	}
}

// 获取文件类型
func getFileType(fileName string) string {
	fileNames := strings.Split(fileName, ".")
	return fileNames[len(fileNames)-1:][0]
}

// 获取文件大小
func getFileSize(size int64) (int64, string) {
	var resSize int64
	var resStr string
	if size > 1024*1024*1024 { // G
		resSize = size / 1024 / 1024 / 1024
		resStr = "GB"
	} else if size > 1024*1024 { // M
		resSize = size / 1024 / 1024
		resStr = "MB"
	} else if size > 1024 { // KB
		resSize = size / 1024
		resStr = "KB"
	} else {
		resSize = size
		resStr = "字节"
	}
	return resSize, resStr
}
