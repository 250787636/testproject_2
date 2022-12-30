package main

import (
	"ky/ssp/crawler/engine"
	"ky/ssp/crawler/persist"
	"ky/ssp/crawler/scheduler"
	"ky/ssp/crawler/zhenai/parser"
)

func main() {
	//engine.SimpleEngine.Run(engine.Request{
	//	Url:        "https://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	e := engine.ConcurrentEngine{
		Secheduler:  &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	//e.Run(engine.Request{
	//	Url:        "https://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
