package pkg_encoding

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math"
)

func DemoGod() {
	fmt.Println("god包二进制传输,在编码器（发送器）和解码器（接受器）之间交换的binary值。一般用于传递远端程序调用（RPC）的参数和结果，如net/rpc包就有提供。")
	//fmt.Println("模拟网络数据的传输接收")
	//asNetwork()
	//fmt.Println("本示例传输一个实现自定义编码和解码方法的值")
	//demoGodEnOrDe()
	fmt.Println("此示例显示如何编码接口值,注册具体类型实现接口")
	demoInterface()

}

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

// 此示例显示了包的基本用法：创建编码器，
// 传输一些值，用解码器接收。
func asNetwork() {
	// 初始化编码器和解码器。 通常是enc和dec
	// 绑定到网络连接和编码器和解码器会
	// 在不同的进程中运行。
	var network bytes.Buffer        // 替代网络连接
	enc := gob.NewEncoder(&network) // 将写入网络。
	dec := gob.NewDecoder(&network) // 将从网络上读取。
	// Encoding（发送）一些值。
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	err = enc.Encode(P{1782, 1841, 1922, "Treehouse"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// Decode（接收）并打印值。
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 2:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
}

// ------------------本示例传输一个实现自定义编码和解码方法的值。
// ------------------本示例传输一个实现自定义编码和解码方法的值。
// ------------------本示例传输一个实现自定义编码和解码方法的值。

// Vector 类型具有未导出的字段，包无法访问。
// 因此，我们编写了一个BinaryMarshal/BinaryUnmarshal方法对来允许我们
// 使用gob包发送和接收类型。 这些接口是
// 在“encoding”包中定义。
// 我们可以等效地使用本地定义的GobEncode/GobDecoder
// 接口。
type Vector struct {
	x, y, z int
}

func (v Vector) MarshalBinary() ([]byte, error) {
	// 一个简单的编码：纯文本。
	var b bytes.Buffer
	fmt.Fprintln(&b, v.x, v.y, v.z)
	return b.Bytes(), nil
}

// UnmarshalBinary 修改接收器，因此必须使用指针接收器。
func (v *Vector) UnmarshalBinary(data []byte) error {
	// 一个简单的编码：纯文本。
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &v.x, &v.y, &v.z)
	return err
}

// 此示例传输实现自定义编码和解码方法的值。
func demoGodEnOrDe() {
	var network bytes.Buffer // 替代（Stand-in）网络。
	// 创建编码器并发送值。
	enc := gob.NewEncoder(&network)
	err := enc.Encode(Vector{3, 4, 5})
	if err != nil {
		log.Fatal("encode:", err)
	}

	// 创建解码器并接收值。
	dec := gob.NewDecoder(&network)
	var v Vector
	err = dec.Decode(&v)
	if err != nil {
		log.Fatal("decode:", err)
	}
	fmt.Println(v)
}

// 此示例显示如何编码接口值。 关键的与常规类型的区别是注册具体类型实现接口。
// 此示例显示如何编码接口值。 关键的与常规类型的区别是注册具体类型实现接口。
// 此示例显示如何编码接口值。 关键的与常规类型的区别是注册具体类型实现接口。

type Point struct {
	X, Y int
}

func (p Point) Hypotenuse() float64 {
	return math.Hypot(float64(p.X), float64(p.Y))
}

type Pythagoras interface {
	Hypotenuse() float64
}

// demoInterface 此示例显示如何编码接口值。 关键的
// 与常规类型的区别是注册具体类型
// 实现接口。
func demoInterface() {
	var network bytes.Buffer // 替代（Stand-in）网络。

	// 我们必须注册编码器和解码器的具体类型（这将是
	// 通常在与编码器不同的机器上）。 在每一端，这告诉了
	// 发送具体类型的引擎实现接口。
	gob.Register(Point{})

	// 创建编码器并发送一些值。
	enc := gob.NewEncoder(&network)
	for i := 1; i <= 3; i++ {
		interfaceEncode(enc, Point{3 * i, 4 * i})
	}

	// 创建解码器并接收一些值。
	dec := gob.NewDecoder(&network)
	for i := 1; i <= 3; i++ {
		result := interfaceDecode(dec)
		fmt.Println(result.Hypotenuse())
	}

}

// interfaceEncode 将接口值编码到编码器中。
func interfaceEncode(enc *gob.Encoder, p Pythagoras) {
	// 除非具体类型，否则编码将失败
	// 注册。 我们在调用函数中注册了。
	// 将指针传递给接口，以便Encode看到（并因此发送）一个值
	// 界面类型。 如果我们直接传递p，它会看到具体的类型。
	// 有关背景，请参阅博客文章“（The Laws of Reflection）反思的法则”。
	err := enc.Encode(&p)
	if err != nil {
		log.Fatal("encode:", err)
	}
}

// interfaceDecode 解码流中的下一个接口值并返回。
func interfaceDecode(dec *gob.Decoder) Pythagoras {
	// 除非线路上的具体类型已经解码，否则解码将失败
	// 注册。 我们在调用函数中注册了。
	var p Pythagoras
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal("decode:", err)
	}
	return p
}
