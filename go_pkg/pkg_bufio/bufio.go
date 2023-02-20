package pkg_bufio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DemoBufio() {
	//input := "1234 5678 1234567901234567890"
	//fmt.Println(splitScanner(input))
	//input := "1,2,3,4,"
	//fmt.Println(commaSplitScanner(input))
	//input := "1111\n22222\n3433333"
	//fmt.Println(lineSplitScanner(input))
	// 人为输入源。
	//input := "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	//wordSplitScanner(input)
	stdoutWriter()
}

// 空格拆分 读 扫描功能
func splitScanner(input string) error {
	// 读取内容 并 创建 scanner
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 创建自定义拆分
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err != nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	// 设置扫描分割
	scanner.Split(split)
	// 验证输入
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
		return err
	}
	return nil
}

// 逗号拆分 读
func commaSplitScanner(input string) error {
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 逗号分割函数
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		// 有一个最终的令牌要交付，可能是空字符串。
		// 在这里返回bufio.ErrFinalToken告诉Scan，在此之后没有更多的标记
		// 但不会触发从扫描本身返回的错误。
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	for scanner.Scan() {
		fmt.Printf("%q\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
		return err
	}
	return nil
}

// 行拆分 读
func lineSplitScanner(input string) error {
	// 标准输入
	// scanner := bufio.NewScanner(os.Stdin)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // println 将添加最后的 \n
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return err
	}
	return nil
}

// 词拆分 读
func wordSplitScanner(input string) error {
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 设置扫描操作的分割功能。
	scanner.Split(bufio.ScanWords)
	// 计算单词
	count := 0
	for scanner.Scan() {
		count++
		fmt.Println(scanner.Text()) // println 将添加最后的 \n
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
		return err
	}
	fmt.Printf("%d\n", count)
	return nil
}

// 写 示例到终端
func stdoutWriter() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")
	w.Flush() // 不要忘记刷新！
}
