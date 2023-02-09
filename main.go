package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

/*
   @Auth: menah3m
   @Desc: 抓取 澎湃新闻首页内容的简单demo
*/

func main() {
	url := "https://www.thepaper.cn/"

	body, err := Fetch(url)
	if err != nil {
		fmt.Println("fetch url err:", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Println("read content err:", err)
	}
	doc.Find("div.index-leftside ").Each(func(i int, selection *goquery.Selection) {
		title := selection.Text()
		fmt.Println(title)
	})
}

func Fetch(url string) ([]byte, error) {

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

	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewEncoder())
	return ioutil.ReadAll(utf8Reader)
}

func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		fmt.Println("fetch err: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func TextReg(regP string, body []byte) [][]byte {
	res := make([][]byte, 1)
	headerRe := regexp.MustCompile(regP)
	matches := headerRe.FindAllSubmatch(body, -1)
	for _, m := range matches {
		res = append(res, m[1])
	}
	return res
}
