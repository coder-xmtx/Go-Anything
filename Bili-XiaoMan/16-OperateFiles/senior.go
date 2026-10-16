package main

import "os"

func main() {

	// OpenFile 操作文件，第二个参数需要传入指定flag（可读/可写/...），第三个参数是权限
	file, _ := os.OpenFile("senior.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o777)

	// 第三个参数权限对 windows 是无效的，针对 linux 和 macos 这类系统

	file.Close()
}
