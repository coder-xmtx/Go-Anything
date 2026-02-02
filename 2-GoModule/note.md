# Go Modules

go.mod 文件是go项目的依赖文件，类似NodeJS中的package.json

```bash
go mod init xxx
```

go 项目的文件树有一个标准，具体查阅文档

- cmd 文件夹存放可执行文件
- internal 文件夹存放私用的函数之类的

main.go文件是这个项目的入口文件
