package pkg_image

import (
	"encoding/base64"
	"fmt"
	"image"
	"log"
	"strings"

	// _ "image/gif"
	// _ "image/png"
	_ "image/jpeg"
)

func DemoImage() {
	fmt.Println("image实现了基本的2D图片库,暂时不深入")
	// 解码JPEG数据。 如果从文件中读取，请创建一个阅读器
	// reader, err := os.Open("testdata/video-001.q50.420.jpeg")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer reader.Close()
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(ImageData))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

	// 计算m的红色，绿色，蓝色和alpha分量的16-bin直方图。
	//
	// 图像的边界不一定从（0,0）开始，因此两个循环开始
	// 在bounds.Min.Y和bounds.Min.X。 首先循环Y和X秒更多
	// 可能导致比X first和Y second更好的内存访问模式。
	var histogram [16][4]int
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			// 颜色的RGBA方法返回[0,65535]范围内的值。
			// 移位12将其减小到[0,15] .g的范围
			histogram[r>>12][0]++
			histogram[g>>12][1]++
			histogram[b>>12][2]++
			histogram[a>>12][3]++
		}
	}

	// 打印结果。
	fmt.Printf("%-14s %6s %6s %6s %6s\n", "bin", "red", "green", "blue", "alpha")
	for i, x := range histogram {
		fmt.Printf("0x%04x-0x%04x: %6d %6d %6d %6d\n", i<<12, (i+1)<<12-1, x[0], x[1], x[2], x[3])
	}
}
