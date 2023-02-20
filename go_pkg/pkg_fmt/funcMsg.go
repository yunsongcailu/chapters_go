package pkg_fmt

const FmtFuncMsg string = `
Printf    格式化字符串写入    标准输出
Fprintf   格式化字符串并写入   w
Sprintf   格式化字符串       返回该字符串
Print     默认格式格式化写入  标准输出
Fprint    默认格式格式化写入  w
Sprint    默认格式格式化，串联所有输出生成并返回一个字符串
Println   默认格式格式化写入   标准输出
Fprintln  默认格式格式化并写入  w
Sprintln  默认格式格式化，串联所有输出生成并返回一个字符串
Errorf    根据format参数生成格式化字符串并返回一个包含该字符串的错误
Scanf     从标准输入扫描文本
Fscanf    从r扫描文本
Sscanf    从字符串str扫描文本
Scan      从标准输入扫描文本
Fscan     从r扫描文本
Sscan     从字符串str扫描文本
Scanln    类似Scan，但会在换行时才停止扫描
Fscanln   类似Fscan，但会在换行时才停止扫描
Sscanln   类似Sscan，但会在换行时才停止扫描
`
