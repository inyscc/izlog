# izlog

运行前需要在 `src/runtime/runtime2.go`文件中添加
```Go
func Goid() int64 { return getg().goid }
```