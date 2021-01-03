package reverseproxy

import (
	"net/url"
	"testing"
)

func TestURL(t *testing.T){
	t.Log(url.PathUnescape("%5B%E5%92%96%E5%95%A1%E8%B1%86%5D.mp3"))
}
