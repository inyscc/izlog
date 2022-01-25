# izlog

将只会打印下面形式的日志：
```log
[2022-01-25 23:05:44.181][ForTest][DEBUG][stu/main.go.main:12][goid:1]去年今日此门中
[2022-01-25 23:05:44.181][ForTest][INFO][stu/main.go.main:13][goid:1]去年今日此门中
[2022-01-25 23:05:44.181][ForTest][ERROR][stu/main.go.main:14][goid:1]去年今日此门中
```


运行前需要在 `src/runtime/runtime2.go`文件中添加
```Go
func Goid() int64 { return getg().goid }
```