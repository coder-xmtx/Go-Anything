# 初级架构

## 1. 初始化项目/模块

```bash
go mod init [模块名称]
```

模块的名称有两种，一种公开，一种私有

- 公开： `github.com/用户名/仓库名`
- 私有： 自己搞着玩的就随便起

## 2. 运行

Golang 可以单文件运行：

```bash
go run main.go
```

也可以整个项目运行，Golang 会查找包名有 main 的 Go 文件

```bash
go run ./20-BasicArchitecture
```

## 3. 多模块协作

类似此文件夹中的 `goods` 和 `order`，可以单独起包名，然后 `main` 去调用

包里的方法和变量开头要大写才能被其他模块调用，即公共化

## 4. 包名一致

同一目录下的包名都要一致

这就保证了有关联的文件就放到一起，没有关联的就别往里面塞垃圾，Go 的设计哲学

注意：子目录就没有这个规则

## 5. 第三方模块

默认从 https://pkg.go.dev 获取 Go Module

```bash
go get 包名
```

外网访问较慢，需要配置国内代理：

```bash
# 启用 Go Modules 功能
go env -w GO111MODULE=on

# 七牛云代理
go env -w GOPROXY=https://goproxy.cn,direct
```

验证配置成功：

```bash
go env GOPROXY
```

【额外】

Go 和 npm 这种不一样，Go的Module会全局缓存到一个文件夹里，不会在项目文件夹里

但是默认是 C 盘，如果想更改默认 cache 目录，在 Windows 下的方法：

- 在环境变量里的Path添加一条路径，比如`D:\GoPath\bin`
- 在环境变量额外写入`GOPATH`，值为`D:\GoPath`

## 6. 依赖治理

没用的模块删除，有用的模块加载

```bash
go mod tidy
```

## 7. 目录结构搭建

```
hello-mod/
├── go.mod
├── cmd/
│   └── hello/
│       └── main.go      # 可执行入口：go run ./cmd/hello
├── internal/
│   └── auth/
│       └── auth.go      # 仅本模块能导入（编译器强制）
└── pkg/
    └── greeter/
        └── greeter.go   # 可被外部模块导入的库代码

```

- `cmd` 文件夹是程序入口
- `internal` 文件夹不能被外部调用，仅有本模块能调用
- `pkg` 公共代码文件夹
