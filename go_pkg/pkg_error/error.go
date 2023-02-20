package pkg_error

import (
	"errors"
	"fmt"
	"time"
)

func DemoError() {
	fmt.Println("Errors 包实现了处理错误的函数")
	if err := oops(); err != nil {
		fmt.Println(err)
	}
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}

// MyError 是一个包含时间和消息的错误实现。
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}
