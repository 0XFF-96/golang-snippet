package main

import (
	"fmt"
	"hash/fnv"
	"math"
	"strings"
)

const (
	SIMILAR_DISTANCE = 3
)

type WordWeight struct {
	Word   string
	Weight float64
}

func SimHashSimilar(srcWordWeighs, dstWordWeights []WordWeight) (distance int, err error) {

	srcFingerPrint, err := simhashFingerPrint(srcWordWeighs)
	if err != nil {
		return
	}
	fmt.Println("srcFingerPrint: ", srcFingerPrint)
	dstFingerPrint, err := simhashFingerPrint(dstWordWeights)
	if err != nil {
		return
	}
	fmt.Println("dstFingerPrint: ", dstFingerPrint)

	distance = hammingDistance(srcFingerPrint, dstFingerPrint)

	return
}

func simhashFingerPrint(wordWeights []WordWeight) (fingerPrint []string, err error) {
	binaryWeights := make([]float64, 32)
	for _, ww := range wordWeights {
		bitHash := strHashBitCode(ww.Word)
		weights := calcWithWeight(bitHash, ww.Weight) //binary每个元素与weight的乘积结果数组
		binaryWeights, err = sliceInnerPlus(binaryWeights, weights)
		//fmt.Printf("ww.Word:%v, bitHash:%v, ww.Weight:%v, binaryWeights: %v\n", ww.Word,bitHash, ww.Weight, binaryWeights)
		if err != nil {
			return
		}
	}
	fingerPrint = make([]string, 0)
	for _, b := range binaryWeights {
		if b > 0 { // bit 1
			fingerPrint = append(fingerPrint, "1")
		} else { // bit 0
			fingerPrint = append(fingerPrint, "0")
		}
	}

	return
}

func strHashBitCode(str string) string {
	h := fnv.New32a()
	h.Write([]byte(str))
	b := int64(h.Sum32())
	return fmt.Sprintf("%032b", b)
}

func calcWithWeight(bitHash string, weight float64) []float64 {
	bitHashs := strings.Split(bitHash, "")
	binarys := make([]float64, 0)

	for _, bit := range bitHashs {
		if bit == "0" {
			binarys = append(binarys, float64(-1)*weight)
		} else {
			binarys = append(binarys, float64(weight))
		}
	}

	return binarys
}

func sliceInnerPlus(arr1, arr2 []float64) (dstArr []float64, err error) {
	dstArr = make([]float64, len(arr1), len(arr1))

	if arr1 == nil || arr2 == nil {
		err = fmt.Errorf("sliceInnerPlus array nil")
		return
	}
	if len(arr1) != len(arr2) {
		err = fmt.Errorf("sliceInnerPlus array Length NOT match, %v != %v", len(arr1), len(arr2))
		return
	}

	for i, v1 := range arr1 {
		dstArr[i] = v1 + arr2[i]
	}

	return
}

func hammingDistance(arr1, arr2 []string) int {
	count := 0
	for i, v1 := range arr1 {
		if v1 != arr2[i] {
			count++
		}
	}
	return count
}


func CosineSimilar(srcWords, dstWords []string) float64 {
	allWordsMap := make(map[string]int, 0)
	for _, word := range srcWords {
		allWordsMap[word] += 1
	}
	for _, word := range dstWords {
		allWordsMap[word] += 1
	}

	// stable the sort
	allWordsSlice := make([]string, 0)
	for word, _ := range allWordsMap {
		allWordsSlice = append(allWordsSlice, word)
	}

	// vector
	// 向量化
	srcVector := make([]int, len(allWordsSlice))
	dstVector := make([]int, len(allWordsSlice))
	for _, word := range srcWords {
		if index := indexOfSlice(allWordsSlice, word); index != -1 {
			srcVector[index] += 1
		}
	}
	for _, word := range dstWords {
		if index := indexOfSlice(allWordsSlice, word); index != -1 {
			dstVector[index] += 1
		}
	}

	// calc cos
	numerator := float64(0)
	srcSq := 0
	dstSq := 0
	for i, srcCount := range srcVector {
		dstCount := dstVector[i]
		numerator += float64(srcCount * dstCount)
		srcSq += srcCount * srcCount
		dstSq += dstCount * dstCount
	}
	denominator := math.Sqrt(float64(srcSq * dstSq))

	return numerator / denominator
}

func indexOfSlice(ss []string, s string) (index int) {
	index = -1
	for k, v := range ss {
		if s == v {
			index = k
			break
		}
	}

	return
}
