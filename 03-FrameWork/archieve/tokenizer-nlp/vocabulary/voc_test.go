package vocabulary

import (
	"testing"

	"github.com/karan/vocabulary"
	"github.com/stretchr/testify/require"
)

func TestVocabulary(t *testing.T){
	// 直接放弃，这里要求的是字段
	voca, err := vocabulary.New(&vocabulary.Config{})
	require.NoError(t, err)
	_ = voca
	// voca.Antonyms()
}
