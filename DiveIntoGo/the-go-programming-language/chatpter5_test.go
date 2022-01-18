package the_go_programming_language

import (
	"fmt"
	"time"
)

// page 139
// breadthFirst calls f for each item in the worklist
// Any items returned by f are added to the worklist
// f is called at most once for each time
func breadFirst(f func(item string) []string, worklist []string){
	seem := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist           // 在两次迭代之间，用两个队列相互保存元素的技巧
		worklist = nil
		for _, item := range items {
			seem[item] = true
			worklist = append(worklist, f(item)...)
		}
	}
}

// The right place for a defer statement that
// releases a resource is immediately after the resource
// has been successfully acquired.

// 如何对一个函数进行延迟计时
// how to invoke a function values
func trace(msg string) func(){
	start := time.Now()
	// do something
	return func(){fmt.Printf("exit %s %s", msg, time.Since(start))}
}
// 调用方法是： trace("msg msg")(), 记得需要两个括号...
// 延迟计算的最好的例子...