package main

const version = "v0.0.4"

// Args flag 参数
type Args struct {
	ShowVersion      bool   // 显示当前版本
	TrimPrefix       bool   // 去掉文件前缀
	DisableOrComment bool   // 不使用注释作为mapping
	Package          string // 覆盖默认包名
	ModelImportPath  string //
	Schema           string // file+mysql, file+tidb
}
