package BenchMark

import (
	"encoding/json"
	"testing"
)

type JSON struct {
	Foo int
	Bar string
	Baz float64
}

func BenchmarkJsonMarshall(b *testing.B) {
	j := JSON{
		Foo: 123,
		Bar: `benchmark`,
		Baz: 123.456,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(&j)
	}
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	bytes := `{"foo": 1, "bar": "my string", bar: 1.123}`
	str := []byte(bytes)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := JSON{}
		_ = json.Unmarshal(str, &j)
	}
}
