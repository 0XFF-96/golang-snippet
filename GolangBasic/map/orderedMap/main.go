package orderedMap

import (
	"reflect"
	"sort"
)

type OrderedMap struct {
	keys []interface{}
	m    map[interface{}]interface{}
}

// 获取键值对的数量
func (omap *OrderedMap) Len() int {
	return len(omap.keys)
}

func (omap *OrderedMap) Less(i, j int) bool {
	// 省略若干条语句
	return true
}

func (omap *OrderedMap) Swap(i, j int) {
	omap.keys[i], omap.keys[j] = omap.keys[j], omap.keys[i]
}

type Keys interface {
	sort.Interface
	Add(k interface{}) bool
	Remove(k interface{}) bool
	Clear()
	Get(index int) interface{}
	GetAll() []interface{}
	Search(k interface{}) (index int, contains bool)
	ElemType() reflect.Type
	CompareFunc() func(interface{}, interface{}) int8
}

type myKeys struct {
	container   []interface{}
	compareFunc func(interface{}, interface{}) int8
	elemType    reflect.Type
}

func (keys *myKeys) Len() int {
	return len(keys.container)
}

//该方法中，比较两个元素值的操作全权交给了compareFunc字段所代表的那个函数
func (keys *myKeys) Less(i, j int) bool {
	return keys.compareFunc(keys.container[i], keys.container[j]) == -1
}

func (keys *myKeys) Swap(i, j int) {
	keys.container[i], keys.container[j] = keys.container[j], keys.container[i]
}

func (keys *myKeys) isAcceptableElem(k interface{}) bool {
	if k == nil {
		return false
	}
	if reflect.TypeOf(k) != keys.elemType {
		return false
	}
	return true
}

func (keys *myKeys) Add(k interface{}) bool {
	ok := keys.isAcceptableElem(k)
	if !ok {
		return false
	}
	keys.container = append(keys.container, k)
	// sort.Sort函数会通过对keys的值的Len、Less和Swap方法的调用来完成排序。
	// 而在Less方法中，通过compareFunc函数对相邻的元素值进行比较的。
	sort.Sort(keys)
	return true
}

func Search(n int, f func(int) bool) int {
	//省略若干条语句
	return 0
}

//func NewKeys ( compareFunc func(interface{}, interface{}) int8, elemType reflect.Type) Keys {
//	//因为只有*myKeys类型的方法集合中才包含了Keys接口类型中声明的所有方法，
//	//所以下面返回的是一个*myKeys类型值，而不是一个myKeys类型值。
//	return &myKeys {
//		container :         make([]interface{}, 0),
//		compareFunc :   compareFunc,
//		elemType :      elemType,
//	}
//}

type myOrderedMap struct {
	keys     Keys
	elemType reflect.Type
	m        map[interface{}]interface{}
}

// key 的规范
// 字段keys的值中的元素值应该都是有序的，应该可以方便地比较它们之间的大小。
// 字段keys的值的元素类型不应该是一个具体的类型，应该可以在运行时再确定它的元素类型。
// 字段keys的值应该可以方便地进行添加元素值、删除元素值以及获取元素值等操作。
// 字段keys的值中的元素值应该可以被依照固定的顺序获取。
// 字段keys的值中的元素值应该能够被自动地排序。
// 字段keys的值总是已排序的，应该能够确定某一个元素值的具体位置。
// 字段keys既然可以在运行时决定它的值的元素类型，那么就可以在运行时获知这个元素类型。
// 字段keys的值中的不同元素值比较

// https://medium.com/swlh/an-ordered-map-in-go-436634692381
