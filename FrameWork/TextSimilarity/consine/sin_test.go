package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimHashSimilar(t *testing.T) {
	distance, err := SimHashSimilar([]WordWeight{
		{
			Word: "a",
			Weight: 1,
		},
		{
			Word: "b",
			Weight: 1,
		},
		{
			Word: "d",
			Weight: 1,
		},
	}, []WordWeight{
		{
			Word: "a",
			Weight: 1,
		},
		{
			Word: "d",
			Weight: 1,
		},
		{
			Word: "f",
			Weight: 1,
		},
	})
	if err != nil {
		t.Errorf("failed: %v", err.Error())
	}

	t.Logf("SimHashSimilar distance: %v", distance)
}


func TestCosine(t *testing.T){
}

func TestCosineSimilar(t *testing.T) {
	//score := CosineSimilar([]string{"a", "b", "d"}, []string{"a", "d", "f"})
	//t.Logf("CosineSimilar score: %v", score)

	score := CosineSimilar([]string{"a", "b", "d"}, []string{"a", "d", "f"})
	t.Logf("CosineSimilar score: %v", score)
	// 人工计算过的结果 OK
	require.Equal(t, score, 2.0/3.0)
}