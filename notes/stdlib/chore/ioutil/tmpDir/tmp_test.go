package tmpDir

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestTmp(t *testing.T) {
	dir, err := ioutil.TempDir("", "123")
	if err != nil {
		panic(err)
	}
	t.Log(dir)

	// 这个表现：dir 路径不存在就会
}

// initializeDir makes the directory and parent directories if the path doesn't exists yet.
func initializeDir(path string) {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		os.MkdirAll(path, 0700)
	}
}
