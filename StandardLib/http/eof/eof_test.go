package eof

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

//func TestGetObject(t *testing.T) {
//	firebaseRoot := New(firebase_url)
//	body, err := firebaseRoot.Get("1")
//	if err != nil {
//		t.Errorf("Error: %s", err)
//	}
//	t.Logf("%q", body)
//}
//
//func TestPushObject(t *testing.T) {
//	firebaseRoot := New(firebase_url)
//	msg := Message{"testing", "1..2..3"}
//	body, err := firebaseRoot.Push("/", msg)
//	if err != nil {
//		t.Errorf("Error: %s", err)
//	}
//	t.Logf("%q", body)
//}


func TestHttpGet(t *testing.T) {
	client := http.Client{}
	client.Do(nil)
}

func TestHttpRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "www.baidu.com", nil)
	require.NoError(t, err)
	_ = req
}