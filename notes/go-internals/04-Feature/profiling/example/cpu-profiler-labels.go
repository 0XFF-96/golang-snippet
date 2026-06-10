package main

import (
	"bytes"
	"context"
	"fmt"
	// "io/ioutil"
	// "os"
	"runtime/pprof"
	"time"
)

func main() {
	// With gccgo, if a profiling signal arrives at the wrong time
	// during traceback, it may crash or hang. See issue #29448.
	//f, err := ioutil.TempFile("", "proftraceback")
	//if err != nil {
	//	// t.Fatalf("TempFile: %v", err)
	//	fmt.Errorf("%+v", err)
	//}
	//// defer os.Remove(f.Name())
	//defer f.Close()

	var prof bytes.Buffer
	pprof.NewProfile("haha")

	pprof.StartCPUProfile(&prof)

	fmt.Println(prof.String())
	fmt.Println("-------")

	defer pprof.StopCPUProfile()

	go work(context.Background(), "alice")
	go work(context.Background(), "bob")

	time.Sleep(5 * time.Second)
}

func work(ctx context.Context, user string) {
	labels := pprof.Labels("user", user)
	pprof.Do(ctx, labels, func(_ context.Context) {
		go backgroundWork()
		directWork()
	})
}

func directWork() {
	for {
	}
}

func backgroundWork() {
	for {
	}
}