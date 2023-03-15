package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/**
* @Description
* @Author lijiesheng
* @Date 2023/3/13 4:20 PM
* golang 解析 html
**/
func main() {
	res, err := http.Get("http://openapi.baidu.com/oauth/2.0/authorize?response_type=code&client_id=wvZNZMVpsyCeI4QzULG3KbGTyRIpe1qQ&redirect_uri=oob&scope=basic,netdisk")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	fmt.Println(res.Body)

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("err")
	}

	htmlString := string(b)
	fmt.Println(htmlString)

	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		panic(err)
	}
	values := make(map[string]string)
	var findInput func(*html.Node)
	findInput = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "input" {
			var key, value string
			for _, attr := range n.Attr {
				if attr.Key == "name" {
					key = attr.Val
				}
				if attr.Key == "value" {
					value = attr.Val
				}
			}
			if key != "" && value != "" {
				values[key] = value
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findInput(c)
		}
	}
	findInput(doc)
	fmt.Printf("%+v\n", values)
}
