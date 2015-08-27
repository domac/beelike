package helper
import (
	"runtime"
	"fmt"
)

//本函数用于捕获Panic并日志输出
func CatchPanic(err *error, sessionID string, funcionName string) {
	if r := recover(); r!=nil {
		buf := make([]byte, 1000)
		runtime.Stack(buf, false)
		if err != nil {
			*err = fmt.Errorf("%v", r)
		}
	}
}
