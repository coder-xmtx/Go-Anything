package greet

import "strings"

// 需要导出外部的函数头必须大写
func Hello(name string) string {
	clean := normalizeName(name)
	return "Hello " + clean
}

// 正常在该文件内部使用
func normalizeName(name string) string {
	n := strings.TrimSpace(name) // 去空格

	if n == "" {
		return "Guest"
	}

	return strings.ToUpper(n)
}
