### README
1. 代码来自于 import "github.com/mwitkow/grpc-proxy/proxy"



### 基本数据结构
1. codec 
n. 编码解码器；多媒体数字信号编解码器
A codec is a device or computer program for encoding or decoding a digital data stream or signal. 
Codec is a portmanteau of coder-decoder or, less commonly, compressor-decompressor.
A codec encodes a data stream or signal for transmission, storage or encryption, 
or decodes it for playback or editing. Codecs are used in videoconferencing, streaming media, and video editing applications.

2. StreamDirector  流转换器/流引导器



3. handler 核心方法

```
func (s *handler) handler(srv interface{}, serverStream grpc.ServerStream) error {}
```
  
### FAQ 

1. 为什么 会有 backend conn 和 client conn 相关的东西？
2. 

