package collect

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"regexp"
)

/*
   @Auth: menah3m
   @Desc:
*/

type Fetcher interface {
	Get(url string) ([]byte, error)
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
