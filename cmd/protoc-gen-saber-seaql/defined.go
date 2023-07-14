package main

const version = "v0.0.3"

// Args flag 参数
type Args struct {
	ShowVersion      bool // 显示当前版本
	DisableOrComment bool // 不使用注释作为mapping
	TrimPrefix       bool // 去掉文件前缀

	Merge    bool   // 合并到一个文件
	Filename string // 合并文件名, 默认 create_table
}
