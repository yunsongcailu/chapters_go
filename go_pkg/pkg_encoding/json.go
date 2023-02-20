package pkg_encoding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func DemoJson() {
	fmt.Println("JSON 的编码和解码")
	// demoJsonMarshal()
	// demoJsonUnmarshal()
	// demoJsonMarshalOrUn()
	// demoDecoderJson()
	// demoDecoderArrayJson()
	// demoJsonRaw()
	demoJsonRawLater()
}

func demoJsonMarshal() {
	fmt.Println("Marshal 返回 v 的 JSON 编码,示例终端输出.")
	type Road struct {
		Name   string
		Number int
	}
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "=", "\t")
	out.WriteTo(os.Stdout)
	fmt.Println("JSON 不能表示循环数据结构，而 Marshal 不处理它们。将循环结构传递给 Marshal 将导致无限递归")
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err = json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

func demoJsonUnmarshal() {
	fmt.Println("Unmarshal 解析 JSON 编码数据并将结果存储在v指向的值中")
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)
	type Blob struct {
		Name  string
		Order string
	}
	var blobs []Blob
	err := json.Unmarshal(jsonBlob, &blobs)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", blobs)
}

type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

// UnmarshalJSON json转string
func (a *Animal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}

	return nil
}

// MarshalJSON string 转 json
func (a Animal) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	default:
		s = "unknown"
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	}

	return json.Marshal(s)
}

func demoJsonMarshalOrUn() {
	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
	var zoo []Animal
	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
		log.Fatal(err)
	}

	census := make(map[Animal]int)
	for _, animal := range zoo {
		census[animal] += 1
	}

	fmt.Printf("Zoo Census:\n* Gophers: %d\n* Zebras:  %d\n* Unknown: %d\n",
		census[Gopher], census[Zebra], census[Unknown])

}

// 本示例使用 Decoder 来解码不同 JSON 值的流
// 本示例使用 Decoder 来解码不同 JSON 值的流
// 本示例使用 Decoder 来解码不同 JSON 值的流

func demoDecoderJson() {
	fmt.Println("本示例使用 Decoder 来解码不同 JSON 值的流")
	const jsonStream = `
		{"Name": "Ed", "Text": "Knock knock."}
		{"Name": "Sam", "Text": "Who's there?"}
		{"Name": "Ed", "Text": "Go fmt."}
		{"Name": "Sam", "Text": "Go fmt who?"}
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}

// 本示例使用 Decoder 解码 JSON 对象的流数组。
// 本示例使用 Decoder 解码 JSON 对象的流数组。
// 本示例使用 Decoder 解码 JSON 对象的流数组。

func demoDecoderArrayJson() {
	fmt.Println("本示例使用 Decoder 解码 JSON 对象的流数组。")
	const jsonStream = `
	[
		{"Name": "Ed", "Text": "Knock knock."},
		{"Name": "Sam", "Text": "Who's there?"},
		{"Name": "Ed", "Text": "Go fmt."},
		{"Name": "Sam", "Text": "Go fmt who?"},
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	]
`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	// 读开括号
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// 而数组包含值
	for dec.More() {
		var m Message
		// 解码数组值（Message）
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v\n", m.Name, m.Text)
	}

	// 阅读结束括号
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
}

// 这个例子使用 RawMessage 在编组期间使用预先计算的 JSON。
// 这个例子使用 RawMessage 在编组期间使用预先计算的 JSON。
// 这个例子使用 RawMessage 在编组期间使用预先计算的 JSON。
func demoJsonRaw() {
	fmt.Println("这个例子使用 RawMessage 在编组期间使用预先计算的 JSON")
	h := json.RawMessage(`{"precomputed": true}`)

	c := struct {
		Header *json.RawMessage `json:"header"`
		Body   string           `json:"body"`
	}{Header: &h, Body: "Hello Gophers!"}

	b, err := json.MarshalIndent(&c, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

// 本示例使用 RawMessage 延迟解析 JSON 消息的一部分。
// 本示例使用 RawMessage 延迟解析 JSON 消息的一部分。
// 本示例使用 RawMessage 延迟解析 JSON 消息的一部分。

func demoJsonRawLater() {
	fmt.Println("本示例使用 RawMessage 延迟解析 JSON 消息的一部分")
	type Color struct {
		Space string
		Point json.RawMessage // 延迟解析直到我们知道color space
	}
	type RGB struct {
		R uint8
		G uint8
		B uint8
	}
	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}

	var j = []byte(`[
		{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
		{"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}
	]`)
	var colors []Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		log.Fatalln("error:", err)
	}

	for _, c := range colors {
		var dst interface{}
		switch c.Space {
		case "RGB":
			dst = new(RGB)
		case "YCbCr":
			dst = new(YCbCr)
		}
		err := json.Unmarshal(c.Point, dst)
		if err != nil {
			log.Fatalln("error:", err)
		}
		fmt.Println(c.Space, dst)
	}
}
