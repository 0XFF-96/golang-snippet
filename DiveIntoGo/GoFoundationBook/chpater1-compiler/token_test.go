package chpater1_compiler

import (
	"fmt"
	"go/scanner"
	"go/token"
	"testing"

)


func parseTokenExample() {
	src := []byte("cos(x) + 2i*sin(x) // euler")

	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}

func TestT(t *testing.T) {
	parseTokenExample()
}
