package jieba

import (
	"fmt"
	"testing"

	"github.com/yanyiwu/gojieba"
)


// CRUD 集合
func TestSplit(t *testing.T){
	var s string
	var words []string
	use_hmm := true

	gojieba.USER_DICT_PATH = "./user.dict.utf8"
	gojieba.STOP_WORDS_PATH = "./stop_word.dict.utf8"
	x := gojieba.NewJieba()

	// 搞明白这里的 Free 是什么意思？
	// 为什么需要 Free
	defer x.Free()
	//s = "我来到北京清华大学"
	//words = x.CutAll(s)
	//fmt.Println(s)
	//fmt.Println("全模式:", strings.Join(words, "/"))
	//
	//words = x.Cut(s, use_hmm)
	//fmt.Println(s)
	//fmt.Println("精确模式:", strings.Join(words, "/"))
	//s = "比特币"
	//words = x.Cut(s, use_hmm)
	//fmt.Println(s)
	//fmt.Println("精确模式:", strings.Join(words, "/"))
	//
	//x.AddWord("比特币")
	//s = "比特币"
	//words = x.Cut(s, use_hmm)
	//fmt.Println(s)
	//fmt.Println("添加词典后,精确模式:", strings.Join(words, "/"))
	//
	//s = "他来到了网易杭研大厦"
	//words = x.Cut(s, use_hmm)
	//fmt.Println(s)
	//fmt.Println("新词识别:", strings.Join(words, "/"))
	s = "小明硕士毕业于中国科学院计算所，后在日本京都大学深造"
	//words = x.CutForSearch(s, use_hmm)
	//// fmt.Println(s)
	//// fmt.Println("搜索引擎模式:", strings.Join(words, "/"))
	//words = x.Cut(s, use_hmm)
	//t.Log("cut, words", words)
	//words = x.CutAll(s)
	//t.Log("cut all:", words)
	//word := x.Tokenize(s, gojieba.SearchMode, true)
	//t.Log("word", len(word))
	//for _, ww := range word {
	//	t.Log("tokenize", ww)
	// }
	//eword := x.Extract(s, 100)
	//t.Log("extract words:", eword)
	//lword := x.ExtractWithWeight(s, 1000)
	//for _, l := range lword {
	//	t.Log("lword", l)
	//}

	t.Log("cut for search words:", words)
	words = x.Cut("我爱蔡徐坤，我喜欢看他的电影和听周杰伦的音乐", use_hmm)
	t.Log("cut, words", words)
	// 把一些单词，添加到相关的库里面。
	// x.AddWord("蔡徐坤")
	words = x.Cut("我爱蔡徐坤，我喜欢看他的电影和听周杰伦的音乐", use_hmm)
	t.Log("cut, words", words)
	words = x.Cut("蔡徐坤，电影听周杰伦音乐", use_hmm)
	t.Log("cut, words", words)
	words = x.CutAll("蔡徐坤，电影听周杰伦音乐")
	t.Log("cut, words", words)
	// gojieba.STOP_WORDS_PATH
	s = "长春市长春药店"
	words = x.Tag(s)
	t.Log("words", words)
	fmt.Println(s)

	words = x.Tag("我喜欢看电影")
	t.Log("tag words", words)

	words = x.Tag("小明硕士毕业于中国科学院计算所，后在日本京都大学深造")
	t.Log(words)
}


func TestStopWord(t *testing.T){
	// gojieba.STOP_WORDS_PATH = "./stop_word.dict.utf8"
	gojieba.USER_DICT_PATH = "./user.dict.utf8"
	x := gojieba.NewJieba()
	text := "小明硕士毕业于中国科学院计算所，后在日本京都大学深造"
	words := x.Tag(text)
	t.Log(words)

	words = x.Cut(text, true)
	t.Log(words)

	words = x.Cut("喂，你好", true)
	t.Log(words)

	words = x.Tag("喂，你好")
	t.Log(words)

	words = x.Tag("喂，你好, 我是蔡徐坤")
	t.Log(words)
}


func TestExtractor(t *testing.T){
	gojieba.STOP_WORDS_PATH = "./stop_word.dict.utf8"
	gojieba.USER_DICT_PATH = "./user.dict.utf8"
	x := gojieba.NewJieba()

	// 停用词库，增加了 "小明, 小红"
	// 只有在 extract 接口，停用词库才能够生效
	words := x.Extract("小明 想 小红 小君 有点水..", -1)
	t.Log(words)
}










