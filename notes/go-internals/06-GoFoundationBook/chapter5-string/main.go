package chapter5_string


type StringHeader struct {
	Data uintptr
	Len int
}

// QA1: 字符串是 GO 语言中重要的书籍结构，其只能被访问而不能被修改和扩容
// 但是可以通过拼接构造出一个新的字符串。
