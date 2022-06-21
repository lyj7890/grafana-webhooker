package logger

import (
	"io"
	"log"
	"os"
)

func init() {
	f, err := os.OpenFile("webhooker.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	defer func() {
		f.Close()
	}()

	// 同时输出到标准输出和文件
	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
