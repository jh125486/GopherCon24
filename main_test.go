package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
)

var sliceSizes = []int{
	1e0,
	1e1,
	1e2,
	1e3,
	1e4,
	1e5,
	1e6,
}

const intersection = 0.20

var testType string

func TestMain(m *testing.M) {
	testType = os.Args[len(os.Args)-1]
	os.Exit(m.Run())
}

func BenchmarkSolution(b *testing.B) {
	var solution func([]int, []int) []int
	switch strings.ToLower(testType) {
	case "1":
		solution = solution1
	case "2":
		solution = solution2
	case "3":
		solution = solution3
	case "4":
		solution = solution4
	default:
		return
	}

	for _, size := range sliceSizes {
		name := fmt.Sprintf("size=%d", size)
		nums1, nums2 := generateSlices(size, intersection)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				solution(nums1, nums2)
			}
		})
	}
}

func generateSlices(n int, x float64) ([]int, []int) {
	a := make([]int, n)
	for i := range n {
		a[i] = i
	}

	common := int(float64(n) * x)
	b := make([]int, n)
	for i := range b {
		if i < common {
			b[i] = i
		} else {
			b[i] = i + n
		}
	}

	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	rand.Shuffle(len(b), func(i, j int) { b[i], b[j] = b[j], b[i] })

	return a, b
}
