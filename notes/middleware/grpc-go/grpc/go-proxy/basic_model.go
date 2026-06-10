package go_proxy

//// Codec defines the interface gRPC uses to encode and decode messages.
//// Note that implementations of this interface must be thread safe;
//// a Codec's methods can be called from concurrent goroutines.
////
//// Deprecated: use encoding.Codec instead.
//type Codec interface {
//	// Marshal returns the wire format of v.
//	Marshal(v interface{}) ([]byte, error)
//	// Unmarshal parses the wire format into v.
//	Unmarshal(data []byte, v interface{}) error
//	// String returns the name of the Codec implementation.  This is unused by
//	// gRPC.
//	String() string
//}
