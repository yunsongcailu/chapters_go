package pkg_html

import (
	"fmt"
	"html"
)

func DemoHtml() {
	fmt.Println("Html包 提供了用于转义和修改 HTML 文本的功能。")
	htmlEscapeString()
	fmt.Println("暂时不深入")
}

func htmlEscapeString() {
	fmt.Println("EscapeString 将特殊字符（如“<”）转义为“&lt”。它只能转义5个这样的字符：<，>，＆，'和“")
	const s = `"Fran & Freddie's Diner" <tasty@example.com>`
	fmt.Printf("EscapeString:%s\n", html.EscapeString(s))
	str := html.UnescapeString(s)
	fmt.Printf("UnescapeString:%s\n", str)
}
