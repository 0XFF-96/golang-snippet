package chapter3_concurrency

import (
	"path/filepath"
)

func ListDirectory(dir string) ([]string,error){
	return nil, nil
}

func ListDirectoryNonBlocking(dir string) <-chan string {
	return nil
}

func walkDir() {
	filepath.Walk("123", nil )
}