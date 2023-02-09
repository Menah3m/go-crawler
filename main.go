package main

import (
	"fmt"
	"github.com/menah3m/go-crawler/collect"
	"time"
)

/*
   @Auth: menah3m
   @Desc: 抓取 澎湃新闻首页内容的简单demo
*/

func main() {
	// url := "https://www.thepaper.cn/"
	// bf := collect.BaseFetch{}
	// body, err := bf.Get(url)
	// if err != nil {
	// 	fmt.Println("fetch url err:", err)
	// }
	//
	// doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	// if err != nil {
	// 	fmt.Println("read content err:", err)
	// }
	// doc.Find("div.index-leftside ").Each(func(i int, selection *goquery.Selection) {
	// 	title := selection.Text()
	// 	fmt.Println(title)
	// })
	url := "https://book.douban.com/subject/1007305/"
	b := collect.BrowserFetch{Timeout: 5 * time.Second}
	body, err := b.Get(url)
	if err != nil {
		fmt.Println("read content failed:", err)
		return
	}
	fmt.Println(string(body))
}
