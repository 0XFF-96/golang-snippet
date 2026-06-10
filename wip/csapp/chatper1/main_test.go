package chatper1

import (
	"database/sql"
	"fmt"
	"testing"
)

// preprocessing phases
// compilation phases
// Assembly
// Linking
// C 语言程序从 i - s - o 三个后缀的变化
// 不同高级语言产生的汇编代码有可能是相同的

//1.1 Information is bits + context
//1.2 Programs are translated other programs into different forms
//1.3 It pays to understand how compilation systems work
//- optimziing program performance
//- understading link-time errors
//- avoiding security holes.
//1.4
func TestHelloWorld(t *testing.T) {
	var x string
	var y int
	fmt.Scanln(&x,&y)
	fmt.Println("Hello World! from golang")
	fmt.Printf("x:%s y:%d", x, y)
}


func TestCaseLog(t *testing.T) {
	err := ShotByID()
	switch err {
	case nil:
	case sql.ErrNoRows:
	default:
		t.Log("default")
	}
}

func ShotByID() error {
	return sql.ErrNoRows
}