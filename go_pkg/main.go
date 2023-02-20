package main

import "2023/go_pkg/pkg_array"

const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)

func main() {
	// fmt.Printf("常量定义:\nb:%d,kb:%d,mb:%d,gb:%d,tb:%d,pb:%d\n", b, kb, mb, gb, tb, pb)
	// 解压缩 pkg_archive.DemoTar()
	// pkg_archive.DemoZip()
	// 缓冲 pkg_bufio.DemoBufio()
	// pkg_builtin.DemoBuiltin()
	// 整数排序切片 pkg_container.DemoHeap()
	// 双向链 pkg_container.DemoList()
	// 环形链 pkg_container.DemoRing()
	// 上下文 pkg_context.DemoContext()
	// 密钥证书 pkg_crypto.DemoCrypto()
	// 数据库 pkg_database.DemoDatabase()
	// debug 暂时不看
	// 编码 pkg_encoding.DemoEncoding()
	// 错误 pkg_error.DemoError()
	// 命令行解析 pkg_flag.DemoFlag()
	// fmt pkg_fmt.DemoFmt()
	// go pkg_go.DemoGo()
	// 字节 pkg_bytes.DemoBytes()
	// 字符串类似字节
	// 哈希 pkg_hash.DemoHash()
	// html pkg_html.DemoHtml()
	// 图像 pkg_image.DemoImage()
	// 数学 pkg_math.DemoMath()
	// 系统文件 pkg_os.DemoOs()
	// 循环 pkg_loop.DemoLoop()
	// 函数 pkg_func.DemoFunc()
	// 数组 pkg_array.DemoArray()
	// 切片 pkg_array.DemoSlice()
	// map
	pkg_array.DemoMap()
}
