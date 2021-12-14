package utils

import (
	"reflect"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		elements []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"worst case",
			args{[]int{5, 4, 3, 2, 1}},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"best case",
			args{[]int{1, 2, 3, 4, 5}},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"test nil argument",
			args{nil},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.elements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
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
func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)
	for i :=0;i<b.N;i++{
		BubbleSort(els)
	}
}
func BenchmarkSort10(b *testing.B) {
	els := getElements(10)
	for i :=0;i<b.N;i++{
		sort.Ints(els)
	}
}
func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElements(1000)
	for i :=0;i<b.N;i++{
		BubbleSort(els)
	}
}

func BenchmarkSort1000(b *testing.B) {
	els := getElements(1000)
	for i :=0;i<b.N;i++{
		sort.Ints(els)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElements(100000)
	for i :=0;i<b.N;i++{
		BubbleSort(els)
	}
}

func BenchmarkSort100000(b *testing.B) {
	els := getElements(100000)
	for i :=0;i<b.N;i++{
		sort.Ints(els)
	}
}