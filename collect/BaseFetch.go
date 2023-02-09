package collect

import (
	"bufio"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

/*
   @Auth: menah3m
   @Desc:
*/

type BaseFetch struct {
}

func (b BaseFetch) Get(url string) ([]byte, error) {

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
