package exercises

import (
	"sort"
	"sync"
)

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func SelectSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		k := i - 1
		for j := i; j < n; j++ {
			if arr[j] < arr[k] {
				k = j
			}
		}
		if k != i-1 {
			arr[k], arr[i-1] = arr[i-1], arr[k]
		}
	}
	return arr
}

func InsertSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		j := i
		k := arr[i]
		for ; j > 0 && arr[j-1] > k; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = k
	}
	return arr
}

func Partition(arr []int, k int) ([]int, int) {
	n := len(arr)
	if k >= n {
		return arr, -1
	}
	p := arr[k]
	if k != 0 {
		arr[k], arr[0] = arr[0], arr[k]
	}

	i, j := 0, n-1
	for i < j {
		for ; i < j && p <= arr[j]; j-- {
		}
		arr[i] = arr[j]

		for ; i < j && arr[i] <= p; i++ {
		}
		arr[j] = arr[i]
	}
	arr[i] = p
	return arr, i
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	_, k := Partition(arr, 0)
	QuickSort(arr[:k])
	QuickSort(arr[k+1:])
	return arr
}

func QuickSort2(arr []int) []int {
	n := len(arr)
	if n <= 100 {
		return InsertSort(arr)
	}

	p := arr[0]
	i := 0
	j := n - 1
LOOP:
	for {
		for ; p <= arr[j]; j-- {
			if i == j {
				break LOOP
			}
		}

		arr[i] = arr[j]
		for ; arr[i] <= p; i++ {
			if i == j {
				break LOOP
			}
		}
		arr[j] = arr[i]
	}
	arr[i] = p

	QuickSort2(arr[:i])
	QuickSort2(arr[i+1:])
	return arr
}

func goQuickSort(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()

	n := len(arr)
	if n <= 100 {
		InsertSort(arr)
		return
	}

	p := arr[0]
	i := 0
	j := n - 1
LOOP:
	for {
		for ; p <= arr[j]; j-- {
			if i == j {
				break LOOP
			}
		}

		arr[i] = arr[j]
		for ; arr[i] <= p; i++ {
			if i == j {
				break LOOP
			}
		}
		arr[j] = arr[i]
	}
	arr[i] = p

	wg.Add(2)
	go goQuickSort(arr[:i], wg)
	go goQuickSort(arr[i+1:], wg)
}

func QuickSort3(arr []int) []int {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go goQuickSort(arr, wg)
	wg.Wait()

	return arr
}

func heapAdjust(arr []int, p int) []int {
	n := len(arr)
	l := (p << 1) + 1
	largest := p
	if arr[l] > arr[largest] {
		largest = l
	}

	if r := l + 1; r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest == p {
		return arr
	}

	arr[p], arr[largest] = arr[largest], arr[p]
	if largest >= n>>1 {
		// largest节点没有子节点
		return arr
	}

	return heapAdjust(arr, largest)
}

func MaxHeap(arr []int) []int {
	n := len(arr)
	for i := (n >> 1) - 1; i >= 0; i-- {
		heapAdjust(arr, i)
	}
	return arr
}

func heapBubble(a []int, pos int) {
	if pos == 0 {
		return
	}
	p := (pos - 1) >> 1
	if a[pos] > a[p] {
		a[pos], a[p] = a[p], a[pos]
		heapBubble(a, p)
	}
}

func MaxHeap2(arr []int) []int {
	n := len(arr)
	for i := 2; i <= n; i++ {
		heapBubble(arr[:i], i-1)
	}
	return arr
}

func HeapSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	// build max heap
	for i := (n >> 1) - 1; i >= 0; i-- {
		heapAdjust(arr, i)
	}
	arr[0], arr[n-1] = arr[n-1], arr[0]

	for i := n - 1; i >= 2; i-- {
		heapAdjust(arr[:i], 0)
		arr[0], arr[i-1] = arr[i-1], arr[0]
	}
	return arr
}

func GoSort(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func GoStableSort(arr []int) []int {
	sort.Stable(sort.IntSlice(arr))
	return arr
}
