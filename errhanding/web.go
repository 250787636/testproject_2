package main

import (
	"kmesh2/utils/logger"
	"ky/ssp/errhanding/filelistingsever"
	"log"
	"net/http"
	// 获取pprof
	_ "net/http/pprof"
	"os"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(handler appHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic:%v", r)
				http.Error(w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		// system error
		err := handler(w, r)
		if err != nil {
			logger.Warn("Error handing reques: %s", err.Error())
			code := http.StatusOK
			if userErr, ok := err.(userError); ok {
				http.Error(w,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}

			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/",
		errWrapper(filelistingsever.HandleFileList))
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}
