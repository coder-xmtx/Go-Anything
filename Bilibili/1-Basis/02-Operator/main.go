package main

import "fmt"

const (
	writePermission = 1 << iota
	readPermission
	executePermission
)

func main() {
	isAdmin := true
	hasWriteAccess := false
	canEdit := isAdmin && hasWriteAccess
	fmt.Println("是否有编辑权限：", canEdit)
}
