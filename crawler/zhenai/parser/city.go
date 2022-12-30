package parser

import (
	"ky/ssp/crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/zhenghun/[^"]+)"`)
)

//<a href="http://album.zhenai.com/u/1774343032" target="_blank">期待有缘人</a>
//<a href="http://album.zhenai.com/u/1677398857" target="_blank">愿得一人心的美好</a>
// 正则匹配 获取城市名与网址
func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		name := string(m[2])
		//result.Items = append(result.Items, "User"+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, url, name)
			},
		})
	}

	cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}
	return result
}
