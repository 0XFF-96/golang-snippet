package the_go_programming_language

import (
	"net/http"
	"testing"
)

// P191

// 194
type database map[string]int
func TestMux(t *testing.T){
	// test muxtilpler
	db := database{"shoes":50}
	mux := http.NewServeMux()
	// 为什么🈴️这个接口的函数？
	// trace this result back to its source code.
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.HandleFunc("/list", db.list)   // 两种函数之间的区别是什么？


	// net/http provides a global ServeMux instace called DefaultServeMux and package-level functions
	// ...
	http.HandleFunc("/list", db.list)
	t.Log(http.ListenAndServe("localhost:8080", nil )) // 绑定并且监听相关端口....
}

func (db database) list(w http.ResponseWriter, req *http.Request){
	//for item, price := range db {
	//	// fmt.Fprintf(w, "%s: %s\n", item, price)
	//}
}