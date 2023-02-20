package pkg_builtin

import "fmt"

func DemoBuiltin() {
	fm := showFuncList()
	length := len(fm)
	for i := 0; i < length; i++ {
		fmt.Printf("%d,%v\n", i, fm[i])
	}

}

var funcMaps map[int]string

func showFuncList() map[int]string {
	fm := make(map[int]string)
	fm[0] = "append:'追加',"
	fm[1] = "cap:'容量,数组等于len',"
	fm[2] = "close:'关闭',"
	fm[3] = "complex:'构造一个虚数',"
	fm[4] = "copy:'复制',"
	fm[5] = "delete:'删除指定key',"
	fm[6] = "imag:'图像内置返回复数虚部',"
	fm[7] = "real:'内置返回复数实部',"
	fm[8] = "len:'返回长度',"
	fm[9] = "make:'实例化参数类型,返回非指针',"
	fm[10] = "new:'实例化参数类型,返回指针',"
	fm[11] = "panic:'终止函数并返回错误',"
	fm[12] = "print:'自举和调试,输出到标准错误',"
	fm[13] = "println:'自举和调试,输出到标准错误,参数输出之间添加空格，输出结束后换行',"
	fm[14] = "recover:'恢复panic',"
	return fm
}
