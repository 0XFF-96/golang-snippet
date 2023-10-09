package go_proxy

import (
	"fmt"

	"google.golang.org/grpc"
)

type rawCodec struct {
	parentCodec grpc.Codec
}

//

// CodecWithParent returns a proxying grpc.Codec with a user
// provided codec as parent.
// This codec is *crucial* to the functioning of the proxy.
// It allows the proxy server to be oblivious
// to the schema of the forwarded messages.
// It basically treats a gRPC message frame as raw bytes.
// However, if the server handler,
// or the client caller are not proxy-internal functions
// it will fall back to trying to decode the message using a fallback codec.

// 这里讲来两个事情
// 1. Codec 很重要，因为他可以对 forwarded messages 进行 oblivious 对处理（grpc 的钩子函数）
// 2. 如果 server handler/client caller 没有找到，就会 fallback 到 fallback codec.


func CodecWithParent(fallback grpc.Codec) grpc.Codec {
	return &rawCodec{fallback}
}

type frame struct {
	payload []byte
}

func (c *rawCodec) Marshal(v interface{}) ([]byte, error) {
	out, ok := v.(*frame)
	if !ok {
		return c.parentCodec.Marshal(v)
	}
	return out.payload, nil

}

func (c *rawCodec) Unmarshal(data []byte, v interface{}) error {
	dst, ok := v.(*frame)
	if !ok {
		return c.parentCodec.Unmarshal(data, v)
	}
	dst.payload = data
	return nil
}

func (c *rawCodec) String() string {
	return fmt.Sprintf("proxy>%s", c.parentCodec.String())
}
