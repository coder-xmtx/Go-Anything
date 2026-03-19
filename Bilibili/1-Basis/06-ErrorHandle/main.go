package main

import (
	"fmt"
	"time"
)

func main() {
	userID := 2
	name, err := queryDatabase(userID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)
}

type BusinessError struct {
	Code int
	Msg  string
	Time time.Time
}

const TimeFmt = "2006-01-02 15:04:05"

func (b *BusinessError) Error() string {
	return fmt.Sprintf("Code:%d, Msg:%s, Time:%v", b.Code, b.Msg, b.Time.Format(TimeFmt))
}

func queryDatabase(userID int) (string, error) {
	if userID == 1 {
		return "admin", nil
	}
	return "", &BusinessError{Code: 404, Msg: "user not found", Time: time.Now()}
}
