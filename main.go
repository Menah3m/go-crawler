package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
   @Auth: menah3m
   @Desc: 抓取 澎湃新闻首页内容的简单demo
*/

func main() {
	url := "https://www.thepaper.cn/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("get url err:%v", err)
		fmt.Println("请求页面失败")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("status code err:%v", err)
		fmt.Println("返回的状态码不为200")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read body err:%v", err)
		fmt.Println("读取返回内容失败")
	}

	fmt.Println("返回结果：")
	fmt.Println("%s", string(body))
}
