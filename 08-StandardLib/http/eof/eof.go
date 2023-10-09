package eof

import (
	"fmt"
	"html"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

//type client struct {
//
//}

//func (c *client) SendRequest(method string, url, path string, body io.Reader)([]byte, error) {
//	// create a request
//	req, err := http.NewRequest(method, url, body)
//	if err != nil {
//		return nil, err
//	}
//
//	// send JSON to firebase
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		return nil, err
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		return nil, fmt.Errorf("Bad HTTP Response: %v", resp.Status)
//	}
//
//	defer resp.Body.Close()
//	b, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	return b, nil
//}


//// Send HTTP Request, return data
//func (f *firebaseRoot) SendRequest(method string, path string, body io.Reader) ([]byte, error) {
//	url := f.BuildURL(path)
//
//	// create a request
//	req, err := http.NewRequest(method, url, body)
//	if err != nil {
//		return nil, err
//	}
//
//	// send JSON to firebase
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		return nil, err
//	}
//
//	if resp.StatusCode != http.StatusOK {
//		return nil, fmt.Errorf("Bad HTTP Response: %v", resp.Status)
//	}
//
//	defer resp.Body.Close()
//	b, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	return b, nil
//}


func client(addr string) {
	get(addr)
	t0 := time.Now()
	for time.Since(t0) < 2*time.Second {
	}
	get(addr)
}

func get(addr string) {
	resp, err := http.Get("http://" + addr)
	if err != nil {
		log.Fatalf("get: %v", err)
	}
	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
}

type timeoutListener struct {
	net.Listener
}

func (l *timeoutListener) Accept() (net.Conn, error) {
	c, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}
	fmt.Printf("new connection accepted\n")
	go func() {
		time.Sleep(1 * time.Second)
		c.Close()
	}()
	return c, nil
}

func server(addr string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	})
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}
	srv := &http.Server{}
	log.Fatalf("serve: %v", srv.Serve(&timeoutListener{l}))
}
