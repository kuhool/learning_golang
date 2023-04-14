package log

import (
	"fmt"
)

func Test() {

	// log.Println("使用log.Println打印：这是一条很普通的日志。")
	// v := "1111111 or 文本"
	// log.Printf("这是一条%s日志。\n", v)
	//Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。
	// log.Fatalln("这是一条会触发fatal的日志。")
	// log.Panicln("这是一条会触发panic的日志。") //打印panic相关信息

	//设置标准logger的输出选项格式
	// log.SetFlags(log.Llongfile | log.LstdFlags)
	// log.Println("这是一条很普通的日志。")

	//其中Prefix函数用来查看标准logger的输出前缀，SetPrefix函数用来设置输出前缀。
	// log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	// log.Println("这是一条很普通的日志。")
	// log.SetPrefix("[pprof]")
	// log.Println("这是一条很普通的日志。")

	// 例如，下面的代码会把日志输出到同目录下的xx.log文件中。
	// logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Println("open log file failed, err:", err)
	// 	return
	// }
	// log.SetOutput(logFile)
	// log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	// log.Println("这是一条很普通的日志。")
	// log.SetPrefix("[小王子]")
	// log.Println("这是一条很普通的日志。")

	//考虑使用第三方日志库
	// logrus、zap

	// var logger = xlog.NewPackageLogger("github.com/yourorg/yourrepo", "yourpackage")

	// xlog.Debug(ctx context.Context, "title", "content")
	// fmt.Println(logger)

	fmt.Println("log.Test")
}
