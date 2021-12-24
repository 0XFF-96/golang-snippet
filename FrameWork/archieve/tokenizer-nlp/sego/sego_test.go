package sego

import (
	"testing"

	"github.com/huichen/sego"
)


// 这个也太水了
func TestSego(t *testing.T){
	//var seer sego.Segmenter
	//
	//seer = sego.NewDictionary()

	var segmenter sego.Segmenter
	segmenter.LoadDictionary("github.com/huichen/sego/data/dictionary.txt")

	sg := segmenter.Segment([]byte("我喜欢看电影"))
	t.Log(sg)
}
