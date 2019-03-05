package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

var (
	logFileName                         = flag.String("log", "cServer.log", "Log file name")
	MetricTardisRequestDetailNum string = "tardis-reaper.tardis-request.platform.%s.media_id.%s.imei.%s.idfa.%s.content_id.%s.content_type.%s.action_type.%s"
)

func main() {
	// 定义一个文件
	fileName := "ll.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}

	// 创建一个日志对象
	debugLog := log.New(logFile, "[Debug]", log.LstdFlags)
	debugLog.Println("A debug message here")

	// 配置一个日志格式的前缀
	debugLog.SetPrefix("[Info]")
	debugLog.Println("A Info Message here ")

	// 配置log的Flag参数
	debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	debugLog.Println("A different prefix")

	fmt.Println("------------------------------------")

	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	// set logfile Stdout
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// write log
	log.Printf("%v\n", MetricTardisRequestDetailNum)
}
