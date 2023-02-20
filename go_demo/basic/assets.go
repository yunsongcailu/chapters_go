package basic

var MenuInfo string = `
----------------------------------
	学生管理系统 V0
	  1:添加学生
	  2:删除学生
	  3:修改学生
	  4:查询学生
	  5:学生列表
	  6:退出系统
----------------------------------
`

type Student struct {
	Num  string `json:"num"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
