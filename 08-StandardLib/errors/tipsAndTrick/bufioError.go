package tipsAndTrick

import (
	"io"
)

//b := bufio.NewWriter(fd)
  //  b.Write(p0[a:b])
  //  b.Write(p1[c:d])
  //  b.Write(p2[e:f])
  //  // and so on
  //  if b.Flush() != nil {
  //          return b.Flush()
  //      }
  //  }

type Writer struct {
	err error
	buf []byte
	n   int
	wr  io.Writer
}

// 来自于 bufio 包的错误处理
func (b *Writer) Write(p []byte) (nn int, err error) {
	for len(p) > b.Available() && b.err == nil {
		// do something
	}
	if b.err != nil {
		return nn, b.err
	}
	// do something
	return nn, nil
}

func (b *Writer) Available() int {
	return 0
}
