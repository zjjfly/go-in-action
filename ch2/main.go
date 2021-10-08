package main

import (
	_ "GoInAction/ch2/matchers" //_会让我们调用这个包中所有的init函数初始化包
	"GoInAction/ch2/search"
	"log"
	"os"
)

//init会在运行main之前执行
func init() {
	//将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("President")
}
