package the_go_programming_language

type Point struct {
	X,Y int64
}

// 指针方法与非指针方法有什么区别？
// 修改变量..
// 能否 access 变量的过程..
func (p Point) Add(q Point) Point{return Point{p.X + q.X, p.Y + q.Y}}
func (p Point) Sub(q Point) Point{return Point{p.X - q.X, p.Y - q.Y}}

type Path []Point

// 可以用这种方式进行解耦...
func (path Path) TranslateBy(offset Point,add bool){
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path{
		// Call either path[i].Add(offset) or path[i].Sub(offset)
		path[i] = op(path[i], offset)
	}
	//fmt.Fprint()
}
// 🈶️什么关系...2

type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
//func(s *IntSet) Has(x int) bool {
//	word, bit := x/64, unit(x%64)
//	return word < len(s.words) && s.words[word]&(1<<bit) != 0
//}

