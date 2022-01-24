package izlog

import (
	"runtime"
	"strings"
	"time"
)

// izlog/util.go.caller:9
func caller() *[]byte {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		var a, b = 0, 0
		for i, v := range file {
			if v == '/' {
				a = b
				b = i
			}
		}
		funcNames := strings.Split(runtime.FuncForPC(pc).Name(), ".")
		ret := make([]byte, 0, 20)
		ret = append(ret, file[a:]...)
		ret = append(ret, funcNames[len(funcNames)-1]...)
		ret = append(ret, *toByte(line)...)
		return &ret
	}
	return &[]byte{}
}

type integer interface {
	~int | ~int8 | ~int16 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint64
}

// t == 0 ? [] : expect
func toByte[T integer](t T) *[]byte {
	ret := make([]byte, 0, 64)
	for t > 0 {
		ret = append(ret, byte(t%10))
		t /= 10
	}
	return &ret
}

// [2006-01-02 15:04:05.000]
func now() string {
	return time.Now().Format("[2006-01-02 15:04:05.000]")
}
