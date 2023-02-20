package pkg_encoding

import "fmt"

func DemoEncoding() {
	fmt.Println("encoding包定义了供其它包使用的可以将数据在字节水平和文本表示之间转换的接口")
	fmt.Println("BinaryMarshaler, BinaryUnmarshaler, TextMarshaler, TextUnmarshaler")
	//DemoAscii85()
	//DemoBase32()
	//DemoBase64()
	//DemoBinary()
	//DemoCsv()
	//DemoGod()
	//DemoHex()
	//DemoJson()
	DemoXml()
}
