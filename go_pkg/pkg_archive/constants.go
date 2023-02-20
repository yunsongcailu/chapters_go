package pkg_archive

var ConstType = `
	TypeReg           = '0'    // 普通文件
    TypeRegA          = '\x00' // 普通文件
    TypeLink          = '1'    // 硬链接
    TypeSymlink       = '2'    // 符号链接
    TypeChar          = '3'    // 字符设备节点
    TypeBlock         = '4'    // 块设备节点
    TypeDir           = '5'    // 目录
    TypeFifo          = '6'    // 先进先出队列节点
    TypeCont          = '7'    // 保留位
    TypeXHeader       = 'x'    // 扩展头
    TypeXGlobalHeader = 'g'    // 全局扩展头
    TypeGNULongName   = 'L'    // 下一个文件记录有个长名字
    TypeGNULongLink   = 'K'    // 下一个文件记录指向一个具有长名字的文件
    TypeGNUSparse     = 'S'    // 稀疏文件`

var Variables = `
	ErrWriteTooLong    = errors.New("archive/tar: write too long")
    ErrFieldTooLong    = errors.New("archive/tar: header field too long")
    ErrWriteAfterClose = errors.New("archive/tar: write after close")
	ErrHeader = errors.New("archive/tar: invalid tar header")
	`
var HeaderStruct = `
	Name       string    // 记录头域的文件名
    Mode       int64     // 权限和模式位
    Uid        int       // 所有者的用户ID
    Gid        int       // 所有者的组ID
    Size       int64     // 字节数（长度）
    ModTime    time.Time // 修改时间
    Typeflag   byte      // 记录头的类型
    Linkname   string    // 链接的目标名
    Uname      string    // 所有者的用户名
    Gname      string    // 所有者的组名
    Devmajor   int64     // 字符设备或块设备的major number
    Devminor   int64     // 字符设备或块设备的minor number
    AccessTime time.Time // 访问时间
    ChangeTime time.Time // 状态改变时间
    Xattrs     map[string]string
`
