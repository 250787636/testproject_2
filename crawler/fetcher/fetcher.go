package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

//"https://www.zhenai.com/zhenghun"

var rareLimiter = time.Tick(100 * time.Millisecond)

// 读取页面数据
func Fetch(url string) ([]byte, error) {
	<-rareLimiter
	newUrl := strings.Replace(url, "http://", "https://", 1)
	req, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36")
	cokie1 := "2518829c-6c42-495d-95c1-790b427a39e5; ec=0Mq5MGn7-1649408924117-bd8c3ffe6abc41667761843; notificationPreAuthorizeSwitch=110461; __channelId=901045%2C0; _abTest_ext13_=0; abt_params=za_m_other%7C905828%7C0%7C0%7C0; __detect__grid__=1; _efmdata=%2FiniLxN%2FNtcx6jLYa7KyHgzEUyup7Kpb1mOVEY%2B0%2B9OAvK%2Bs2fUHRtuvT4w%2F%2BwSYFfcDmjeTmRadG4WRcDPHUnVwf8yZ9ep%2B5fqud2eRMhI%3D; _exid=%2BQe4Y9e7Esr5ENBPyzxn%2B1ig8dJ7M0sgdS%2BpNWe4hwn6cFgtHhTAR8OeOmPL13hlrsQF9sTFhAGpu5V%2B9PMspQ%3D%3D; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1649401810,1650351292,1650441027; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1650441027"
	req.Header.Add("cookie", cokie1)
	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// 解码页面数据
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
