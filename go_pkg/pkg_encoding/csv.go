package pkg_encoding

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func DemoCsv() {
	fmt.Println("csv读写逗号分隔值（csv）的文件")
	//demoCsvRead()
	//demoCsvReadAll()
	//demoCsvConfRead()

	//demoCsvWrite()
	//demoCscWriteAll()
	demoCsvWriteFile()
}

func demoCsvRead() {
	fmt.Println("csv read 示例:逗号分割,换行结束")
	in := `first_name,last_name,username,"Rob","Pike",rob
			Ken,Thompson,ken
			"Robert","Griesemer","gri"`
	r := csv.NewReader(strings.NewReader(in))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}

func demoCsvReadAll() {
	fmt.Println("csv read all:ReadAll 从 r 读取所有剩余的记录。每条记录都是一片田地。成功的调用返回 err == nil，而不是 err == io.EOF。由于 ReadAll 被定义为读取直到 EOF，因此它不会将文件末尾视为要报告的错误")
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
}

func demoCsvConfRead() {
	fmt.Println("csv read 示例:自定义符号,分号分割,换行一组,井号跳过")
	in := `first_name;last_name;username
"Rob";"Pike";rob
# lines beginning with a # character are ignored
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`
	r := csv.NewReader(strings.NewReader(in))
	r.Comma = ';'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
}

func demoCsvWrite() {
	fmt.Println("csv write 示例:")
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func demoCscWriteAll() {
	fmt.Println("WriteAll 使用 Write 写入多个 CSV 记录，然后调用 Flush。")
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

func demoCsvWriteFile() {
	file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	defer file.Close()
	// 写入UTF-8 BOM，防止中文乱码
	file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)
	w.Write([]string{"这个是啥子", "我了个擦", "发发发"})
	// 写文件需要flush，不然缓存满了，后面的就写不进去了，只会写一部分
	w.Flush()
}
