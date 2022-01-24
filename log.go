package izlog

import (
	"io"
	"runtime"
	"sync"
)

type Level uint8

const (
	TRACE Level = iota
	DEBUG
	INFO
	ERROR
)

type logger struct {
	io.Writer
	Name string
}

var pool = &sync.Pool{
	New: func() any {
		ret := make([]byte, 0, 200)
		return &ret
	},
}

func (l *logger) write(lvl string, bs string) (n int, err error) {
	b := pool.Get().(*[]byte)
	*b = append(*b, now()...)
	*b = append(*b, lvl...)
	*b = append(*b, l.Name...)
	*b = append(*b, *caller()...)
	*b = append(*b, "[goid:"...)
	*b = append(*b, *toByte(runtime.Goid())...)
	*b = append(*b, ']')
	*b = append(*b, bs...)
	*b = append(*b, '\n')
	n, err = l.Write(*b)
	pool.Put(b)
	return
}

func New(name string, w io.Writer) (l *logger) {
	if w == nil {
		return &logger{
			Writer: w,
			Name:   name,
		}
	}
	return &logger{
		Writer: w,
		Name:   name,
	}
}
