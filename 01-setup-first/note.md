# Setup First

## 1. 安装

前往 [Golang](https://golang.google.cn/) 官网下载安装包

安装的时候直接安装在默认位置即可

## 2. 配置

博主用 VS Code 进行学习，在插件市场下载 Go 的插件

## 3. 主函数

Go 和 C 很类似的一个点，都是用 main.go 来作为主程序入口

## 4. 初始化

打开文件夹后，要先初始化 Go 的配置，在终端执行以下命令

```bash
go mod init
```

## 5. 导入包

`fmt` 是 Go 的标准输入输出流包

```go
// 一个包
import "fmt"

// 多个包
import (
    "fmt"
    "math"
    "strings"
)
```
