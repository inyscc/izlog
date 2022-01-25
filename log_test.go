package izlog_test

import (
	"os"
	"runtime"
	"testing"

	"github.com/inyscc/izlog"
)

var logger = izlog.New("Izlog", os.Stdout)

func BenchmarkIzlog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for range [5]struct{}{} {
			logger.Info("去年今日此门中，人面桃花相映红。人面不知何处去，桃花依旧笑春风。去年今日此门中，人面桃花相映红。人面不知何处去，桃花依旧笑春风。去年今日此门中，人面桃花相映红。人面不知何处去，桃花依旧笑春风。去年今日此门中，人面桃花相映红。人面不知何处去，桃花依旧笑春风。去年今日此门中，人面桃花相映红。人面不知何处去，桃花依旧笑春风。")
		}
	}
}

func TestCaller(t *testing.T) {
	t.Log(runtime.Caller(1))
}
