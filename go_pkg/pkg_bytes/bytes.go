package pkg_bytes

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
)

func DemoBytes() {
	fmt.Println("bytes包实现了操作[]byte的常用函数。本包的函数和strings包的函数相当类似")
	// 字典顺序比较 字节数组
	// bytesCompare()
	// 二进制搜索查找匹配的字节片
	// bytesSearch()
	// 是否包含
	// bytesSublice()
	// ContainsAny 报告字符中的任何 UTF-8 编码的 Unicode 代码点是否在 b 中。
	// bytesContainsAny()
	// ContainsRune Unicode 代码点 r 是否在 b 之内
	// bytesContainsRune()
	// Count 计算s中不重叠实例的数量。如果 sep 为空片段，则 Count 返回1 + s中的 Unicode 代码点数。
	// bytesCount()
	// Equal 返回一个布尔值，报告 a 和 b 是否是相同的长度并且包含相同的字节。零参数相当于一个空片。
	// bytesEqual()
	// EqualFold 判断两个utf-8编码切片（将unicode大写、小写、标题三种格式字符视为相同）是否相同
	// bytesEqualFold()
	// Fields 一个或多个连续空格 分割切片，如果 s 仅包含空格，则返回 s 的子片段或空列表
	// bytesFields()
	// FieldsFunc 将s解释为 UTF-8 编码的 Unicode 代码点序列,每次满足 f(c) 的代码点 c 运行时分割片 s 并返回s的一个子片段
	// f(c) 返回 true一次 则分割一次
	// byteFieldsFunc()
	// HasPrefix测试字节片是否以前缀开头
	// byteHasPrefix()
	// HasSuffix 测试字节片段是否以后缀结尾
	// byteHasSuffix()
	// 索引 返回 s 中第一个 sep 实例的索引，如果 s 中不存在 sep，则返回-1。
	// byteIndex()
	// LastIndex 返回 s 中最后一个 sep 实例的索引，如果 sep 中不存在 s，则返回-1。
	// byteLastIndex()
	// IndexAny 将 s 解释为 UTF-8 编码的 Unicode 代码点序列。
	// 它返回字符中任何 Unicode 代码点的 s 中第一次出现的字节索引。 如果字符为空或者没有共同的代码点，则返回-1。
	// byteIndexAny()
	// IndexByte 返回 s 的第一个c实例的索引，如果 c 不存在于 s 中，则返回 -1。
	// byteIndexByte()
	// IndexFunc 将 s 解释为一系列UTF-8编码的Unicode代码点。它返回满足 f（c） 的第一个 Unicode 代码点的 s
	// 中的字节索引，否则返回 -1。
	// byteIndexFunc()
	// IndexRune将 s 解释为一系列UTF-8编码的Unicode代码点。它返回给定符文的 s 中第一个出现的字节索引。
	// byteIndexRune()
	// join 连接 s的元素以创建一个新的字节片。分隔符 sep 放置在生成的切片中的元素之间。
	// byteJoin()
	// byteMap Map 根据映射函数返回字节切片s的所有字符修改后的副本
	// byteMap()
	// byteRepeat 重复返回由 b 的计数副本组成的新字节片段
	// byteRepeat()
	// byteReplace Replace替换 将返回 slice 的一个副本，其中前 n 个非重叠的旧实例将被 new 替换
	// byteReplace()
	// 将切片分割成由 sep 分隔的所有子切片，并返回这些分隔符之间的一部分子切片
	// byteSplit()
	// byteTitle 大小写
	// byteTitle()
	// 去掉空格
	// byteTrim()
	// buffer缓冲
	byteBuffer()
}

// buffer缓冲
func byteBuffer() {
	fmt.Println("Buffer 的零值是一个准备使用的空缓冲区。缓冲区再做专题")
	var b bytes.Buffer // 缓冲区不需要初始化。
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!")
	b.WriteTo(os.Stdout)

	// 缓冲区可以将字符串或[]字节转换为io.Reader。
	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec)
}

// byteTrim 去掉空格
func byteTrim() {
	fmt.Println("trim 以 cutset 切割 前后")
	fmt.Printf("[%q]", bytes.Trim([]byte(" !!! Achtung! Achtung! !!! "), "! "))
	fmt.Println("TrimFunc,TrimLeft,TrimLeftFunc,TrimPrefix,TrimRight,TrimRightFunc,")
	var b = []byte("Goodbye,, world!")
	b = bytes.TrimPrefix(b, []byte("Goodbye,"))
	b = bytes.TrimPrefix(b, []byte("See ya,"))
	fmt.Printf("Hello%s", b)
	fmt.Println("TrimSpace 所有前导和尾随空格")
	fmt.Printf("%s", bytes.TrimSpace([]byte(" \t\n a lone gopher \n\t\r\n")))
	fmt.Println("TrimSuffix 在没有提供尾随后缀字符串的情况下返回 s 。如果 s 不以后缀结尾，则 s 不变。")
	b = []byte("Hello, goodbye, etc!")
	b = bytes.TrimSuffix(b, []byte("goodbye, etc!"))
	b = bytes.TrimSuffix(b, []byte("gopher"))
	b = append(b, bytes.TrimSuffix([]byte("world!"), []byte("x!"))...)
	os.Stdout.Write(b)
}

// byteTitle 大小写
func byteTitle() {
	fmt.Println("首字母大写title 已弃用")
	fmt.Printf("%s", bytes.Title([]byte("her royal highness")))
	fmt.Println("ToLower 返回所有 Unicode 字母映射为小写字节片段的副本。 ")
	fmt.Printf("%s", bytes.ToLower([]byte("Gopher")))
	fmt.Println("ToLowerSpecial 返回所有 Unicode 字母映射到小写字节的字节切片的副本，优先考虑特殊的外壳规则。 ")
	str := []byte("AHOJ VÝVOJÁRİ GOLANG")
	toTitle := bytes.ToLowerSpecial(unicode.AzeriCase, str)
	fmt.Println("Original : " + string(str))
	fmt.Println("ToLower : " + string(toTitle))
	fmt.Printf("%s\n", bytes.ToTitle([]byte("loud noises")))
	fmt.Printf("%s\n", bytes.ToTitle([]byte("хлеб")))
	fmt.Println(" ToUpper 返回字节切片 s 的副本，并将所有 Unicode 字母映射为大写字母。 ")
	fmt.Printf("%s", bytes.ToUpper([]byte("Gopher")))
}

// byteSplit 将切片分割成由 sep 分隔的所有子切片，并返回这些分隔符之间的一部分子切片
func byteSplit() {
	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a ")))
	fmt.Printf("%q\n", bytes.Split([]byte(" xyz "), []byte("")))
	fmt.Printf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins")))
	fmt.Println("SplitAfter 在 sep 的每个实例之后切片到所有 sublices 中，并返回这些 sublices 的一部分。")
	fmt.Printf("%q\n", bytes.SplitAfter([]byte("a,b,c"), []byte(",")))
	fmt.Println("SplitAfterN 将 s 分割成子项，并返回这些子项的一部分")
	fmt.Printf("%q\n", bytes.SplitAfterN([]byte("a,b,c"), []byte(","), 2))
	fmt.Println("将 SplitN 切片成由 sep 分隔的子片段，并返回这些分隔片之间的一部分子片段")
	fmt.Printf("%q\n", bytes.SplitN([]byte("a,b,c"), []byte(","), 2))
	z := bytes.SplitN([]byte("a,b,c"), []byte(","), 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)
}

// byteReplace Replace替换 将返回 slice 的一个副本，其中前 n 个非重叠的旧实例将被 new 替换
func byteReplace() {
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("ky"), 2))
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("oink"), []byte("moo"), -1))
}

// byteRepeat 重复返回由 b 的计数副本组成的新字节片段
func byteRepeat() {
	fmt.Printf("%s", bytes.Repeat([]byte("na"), 2))
}

// byteMap Map 根据映射函数返回字节切片s的所有字符修改后的副本
func byteMap() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Printf("%s", bytes.Map(rot13, []byte("'Twas brillig and the slithy gopher...")))
}

// bytesCompare 字典顺序比较 字节数组
func bytesCompare() {
	fmt.Println("Compare 按字典顺序返回一个比较两个字节片段的整数")
	fmt.Println("如果a == b，结果将为0，如果 a <b 则返回-1，如果 a> b 则返回+1。零参数相当于一个空片")
	// 通过将比较结果与零比较来解释比较结果。
	var a, b []byte
	a = []byte("aaa")
	b = []byte("bbb")
	if bytes.Compare(a, b) < 0 {
		// a 小于 b
		fmt.Printf("a:%v < b:%v\n", string(a), string(b))
	}
	b = []byte("bbb")
	if bytes.Compare(a, b) <= 0 {
		// a 小于或等于 b
		fmt.Printf("a:%v <= b:%v\n", string(a), string(b))
	}
	a = []byte("ccc")
	if bytes.Compare(a, b) > 0 {
		// a 大于 b
		fmt.Printf("a:%v > b:%v\n", string(a), string(b))
	}
	b = []byte("ccc")
	if bytes.Compare(a, b) >= 0 {
		// a 大于或等于 b
		fmt.Printf("a:%v >= b:%v\n", string(a), string(b))
	}

	// 为了比较平等，优先选择等于比较。
	if bytes.Equal(a, b) {
		// a 等于 b
		fmt.Printf("a:%v = b:%v\n", string(a), string(b))
	}
	b = []byte("ddd")
	if !bytes.Equal(a, b) {
		// a 不等于 b
		fmt.Printf("a:%v != b:%v\n", string(a), string(b))
	}
}

// bytesSearch 二进制搜索查找匹配的字节片
func bytesSearch() {
	var needle []byte
	var haystack [][]byte // 假设排序
	i := sort.Search(len(haystack), func(i int) bool {
		// 返回 haystack[i] >= needle。
		return bytes.Compare(haystack[i], needle) >= 0
	})
	if i < len(haystack) && bytes.Equal(haystack[i], needle) {
		// 找到了
		fmt.Println("找到了")
		return
	}
	fmt.Println("没找到")
	return
}

// bytesSublice 是否包含 sublice 是否在 b 之内
func bytesSublice() {
	fmt.Println("是否包含 sublice 是否在 b 之内")
	b := []byte("seafood")
	sublice := []byte("foo")
	fmt.Printf("b:%v,sublice:%v,isContains:%v\n", string(b), string(sublice), bytes.Contains(b, sublice))
	sublice = []byte("bar")
	fmt.Printf("b:%v,sublice:%v,isContains:%v\n", string(b), string(sublice), bytes.Contains(b, sublice))
	sublice = []byte("")
	fmt.Printf("b:%v,sublice:%v,isContains:%v\n", string(b), string(sublice), bytes.Contains(b, sublice))
	b = []byte("")
	fmt.Printf("b:%v,sublice:%v,isContains:%v\n", string(b), string(sublice), bytes.Contains(b, sublice))
}

// bytesContainsAny 字符中的任何 UTF-8 编码的 Unicode 代码点是否在 b 中。
func bytesContainsAny() {
	fmt.Println("字符中的任何 UTF-8 编码的 Unicode 代码点是否在 b 中。")
	b := []byte("I like seafood.")
	chars := "fÄo!"
	fmt.Printf("b:%v,chars:%v,isContainsAny utf8 chars:%v\n", string(b), chars, bytes.ContainsAny(b, chars))
	chars = "去是伟大的."
	fmt.Printf("b:%v,chars:%v,isContainsAny utf8 chars:%v\n", string(b), chars, bytes.ContainsAny(b, chars))
	chars = ""
	fmt.Printf("b:%v,chars:%v,isContainsAny utf8 chars:%v\n", string(b), chars, bytes.ContainsAny(b, chars))
	b = []byte("")
	fmt.Printf("b:%v,chars:%v,isContainsAny utf8 chars:%v\n", string(b), chars, bytes.ContainsAny(b, chars))
}

// bytesContainsRune Unicode 代码点 r 是否在 b 之内。
func bytesContainsRune() {
	fmt.Println("Unicode 代码点 r 是否在 b 之内")
	b := []byte("I like seafood.")
	r := 'f'
	fmt.Printf("b:%v,r:%v,isContainsRune:%v\n", string(b), string(r), bytes.ContainsRune(b, r))
	r = 'ö'
	fmt.Printf("b:%v,r:%v,isContainsRune:%v\n", string(b), string(r), bytes.ContainsRune(b, r))
	b = []byte("去是伟大的!")
	r = '大'
	fmt.Printf("b:%v,r:%v,isContainsRune:%v\n", string(b), string(r), bytes.ContainsRune(b, r))
	r = '!'
	fmt.Printf("b:%v,r:%v,isContainsRune:%v\n", string(b), string(r), bytes.ContainsRune(b, r))
	r = '@'
	b = []byte("")
	fmt.Printf("b:%v,r:%v,isContainsRune:%v\n", string(b), string(r), bytes.ContainsRune(b, r))
}

// Count 计算s中不重叠实例的数量。如果 sep 为空片段，则 Count 返回1 + s中的 Unicode 代码点数。
func bytesCount() {
	fmt.Println("Count 计算s中不重叠实例的数量。如果 sep 为空片段，则 Count 返回1 + s中的 Unicode 代码点数。")
	s := []byte("cheese")
	sep := []byte("e")
	fmt.Printf("s:%v,sep:%v,count:%d\n", string(s), string(sep), bytes.Count(s, sep))
	sep = []byte("")
	fmt.Printf("s:%v,sep:%v,count:%d\n", string(s), string(sep), bytes.Count(s, sep))
	s = []byte("five")
	fmt.Printf("s:%v,sep:%v,count:%d\n", string(s), string(sep), bytes.Count(s, sep))
}

// bytesEqual Equal 返回一个布尔值，报告 a 和 b 是否是相同的长度并且包含相同的字节。零参数相当于一个空片。
func bytesEqual() {
	fmt.Println("Equal a 和 b 是否是相同的长度并且包含相同的字节。零参数相当于一个空片")
	a := []byte("Go")
	b := []byte("Go")
	fmt.Printf("a:%v,b:%v,equal:%v\n", string(a), string(b), bytes.Equal(a, b))
	b = []byte("C++")
	fmt.Printf("a:%v,b:%v,equal:%v\n", string(a), string(b), bytes.Equal(a, b))
}

// bytesEqualFold EqualFold 判断两个utf-8编码切片（将unicode大写、小写、标题三种格式字符视为相同）是否相同
func bytesEqualFold() {
	var s []byte
	var t []byte
	fmt.Printf("s:%v,t:%v,equalFold:%v\n", string(s), string(t), bytes.EqualFold(s, t))

	s = []byte{'A'}
	t = []byte("A")
	fmt.Printf("s:%v,t:%v,equalFold:%v\n", string(s), string(t), bytes.EqualFold(s, t))
	t = []byte{'a'}
	fmt.Printf("s:%v,t:%v,equalFold:%v\n", string(s), string(t), bytes.EqualFold(s, t))
	s = []byte{'B'}
	fmt.Printf("s:%v,t:%v,equalFold:%v\n", string(s), string(t), bytes.EqualFold(s, t))
	s = []byte{}
	t = nil
	fmt.Printf("s:%v,t:%v,equalFold:%v\n", string(s), string(t), bytes.EqualFold(s, t))
}

// bytesFields Fields 一个或多个连续空格 分割切片，如果 s 仅包含空格，则返回 s 的子片段或空列表
func bytesFields() {
	fmt.Printf("Fields are: %q", bytes.Fields([]byte("  foo bar  baz   ")))
}

// byteFieldsFunc FieldsFunc 将s解释为 UTF-8 编码的 Unicode 代码点序列,每次满足 f(c) 的代码点 c 运行时分割片 s 并返回s的一个子片段
// f(c) 返回 true一次 则分割一次
func byteFieldsFunc() {
	fmt.Println("f(c) 返回 true一次 则分割一次")
	f := func(c rune) bool {
		a := !unicode.IsLetter(c)
		b := !unicode.IsNumber(c)
		d := a && b
		fmt.Printf("c:%v,a:%v,b:%v,a&&b:%v\n", string(c), a, b, d)
		// 字母和数字
		return d
	}
	fmt.Printf("返回 字母和数字 Fields are: %q", bytes.FieldsFunc([]byte("  foo1;bar2,baz3..., add1 "), f))
}

// byteHasPrefix HasPrefix测试字节片是否以前缀开头
func byteHasPrefix() {
	fmt.Println("HasPrefix测试字节片是否以前缀开头")
	s := []byte("Gopher")
	p := []byte("Go")
	fmt.Printf("s:%v,p:%v,hasPrefix:%v\n", string(s), string(p), bytes.HasPrefix(s, p))
	p = []byte("C")
	fmt.Printf("s:%v,p:%v,hasPrefix:%v\n", string(s), string(p), bytes.HasPrefix(s, p))
	p = []byte("")
	fmt.Printf("s:%v,p:%v,hasPrefix:%v\n", string(s), string(p), bytes.HasPrefix(s, p))
}

// byteHasSuffix HasSuffix 测试字节片段是否以后缀结尾
func byteHasSuffix() {
	s := []byte("Amigo")
	suf := []byte("go")
	fmt.Printf("s:%v,suffix:%v,hasSuffix:%v\n", string(s), string(suf), bytes.HasSuffix(s, suf))
	suf = []byte("O")
	fmt.Printf("s:%v,suffix:%v,hasSuffix:%v\n", string(s), string(suf), bytes.HasSuffix(s, suf))
	suf = []byte("")
	fmt.Printf("s:%v,suffix:%v,hasSuffix:%v\n", string(s), string(suf), bytes.HasSuffix(s, suf))

}

// byteIndex 索引 返回 s 中第一个 sep 实例的索引，如果 s 中不存在 sep，则返回-1。
func byteIndex() {
	s := []byte("chicken")
	sep := []byte("ken")
	fmt.Printf("s:%s,sep:%s,sep index in s:%d\n", string(s), string(sep), bytes.Index(s, sep))
	sep = []byte("dmr")
	fmt.Printf("s:%s,sep:%s,sep index in s:%d\n", string(s), string(sep), bytes.Index(s, sep))
}

// byteLastIndex LastIndex 返回 s 中最后一个 sep 实例的索引，如果 sep 中不存在 s，则返回-1。
func byteLastIndex() {
	fmt.Println("lastIndexAny ,lastIndexFunc,lastIndexByte 类似index")
	fmt.Println(bytes.Index([]byte("go gopher"), []byte("go")))
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("go")))
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("rodent")))
}

// byteIndexAny IndexAny 将 s 解释为 UTF-8 编码的 Unicode 代码点序列。
// 它返回字符中任何 Unicode 代码点的 s 中第一次出现的字节索引。 如果字符为空或者没有共同的代码点，则返回-1。
func byteIndexAny() {
	s := []byte("chicken")
	chars := "aeiouy"
	fmt.Printf("s:%s,chars:%s,indexAny chars in s:%d\n", string(s), chars, bytes.IndexAny(s, chars))
	s = []byte("crwth")
	fmt.Printf("s:%s,chars:%s,indexAny chars in s:%d\n", string(s), chars, bytes.IndexAny(s, chars))
}

// byteIndexByte IndexByte 返回 s 的第一个c实例的索引，如果 c 不存在于 s 中，则返回 -1。
func byteIndexByte() {
	s := []byte("chicken")
	c := byte('k')
	fmt.Printf("s:%v,c:%v,indexByte:%d\n", string(s), string(c), bytes.IndexByte(s, c))
	c = byte('g')
	fmt.Printf("s:%v,c:%v,indexByte:%d\n", string(s), string(c), bytes.IndexByte(s, c))
}

// byteIndexFunc IndexFunc 将 s 解释为一系列UTF-8编码的Unicode代码点。它返回满足 f（c） 的第一个 Unicode 代码点的 s
// 中的字节索引，否则返回 -1。
func byteIndexFunc() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Printf("string:Hello, 世界 汉字的index:\n")
	fmt.Println(bytes.IndexFunc([]byte("Hello, 世界"), f))
	fmt.Println(bytes.IndexFunc([]byte("Hello, world"), f))
}

// byteIndexRune IndexRune将 s 解释为一系列UTF-8编码的Unicode代码点。它返回给定符文的 s 中第一个出现的字节索引。
func byteIndexRune() {
	fmt.Println("IndexRune")
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'k'))
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'd'))
}

// byteJoin join 连接 s的元素以创建一个新的字节片。分隔符 sep 放置在生成的切片中的元素之间。
func byteJoin() {
	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("逗号连接 %s", bytes.Join(s, []byte(", ")))
}
