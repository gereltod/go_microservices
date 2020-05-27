package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	// Initialization
	els := []int{11, 29, 103, 4, 20, 17}

	// Execution:
	Sort(els)

	// Validation:
	assert.NotNil(t, els)
	assert.EqualValues(t, 6, len(els))
	assert.EqualValues(t, 4, els[0])
	assert.EqualValues(t, 11, els[1])
	assert.EqualValues(t, 17, els[2])
	assert.EqualValues(t, 20, els[3])
	assert.EqualValues(t, 29, els[4])
	assert.EqualValues(t, 103, els[5])
}

func TestBubbleSortBestCase(t *testing.T) {
	// Initialization
	els := []int{4, 11, 17, 20, 29, 103}

	// Execution:
	Sort(els)

	// Validation:
	assert.NotNil(t, els)
	assert.EqualValues(t, 6, len(els))
	assert.EqualValues(t, 4, els[0])
	assert.EqualValues(t, 11, els[1])
	assert.EqualValues(t, 17, els[2])
	assert.EqualValues(t, 20, els[3])
	assert.EqualValues(t, 29, els[4])
	assert.EqualValues(t, 103, els[5])
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	els := getElements(5)
	assert.NotNil(t, els)
	assert.EqualValues(t, 6, len(els))
	assert.EqualValues(t, 4, els[0])
	assert.EqualValues(t, 11, els[1])
	assert.EqualValues(t, 17, els[2])
	assert.EqualValues(t, 20, els[3])
	assert.EqualValues(t, 29, els[4])
	assert.EqualValues(t, 103, els[5])

}

func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}


func BenchmarkSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkSort1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}