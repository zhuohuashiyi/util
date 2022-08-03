package util

func QuickSort(arr *[]int, begin int, end int) {
	if begin >= end {
		return 
	}
	var k, kValue int
	k = begin
	kValue = (*arr)[begin]
	for i := begin; i <= end; i++ {
		if (*arr)[i] < kValue {
			k++
			if k != i {
				(*arr)[k], (*arr)[i] = (*arr)[i], (*arr)[k]
			}
		}

	}
	(*arr)[begin] = (*arr)[k]
	(*arr)[k] = kValue
	QuickSort(arr, begin, k-1)
	QuickSort(arr, k+1, end)
}

func HeapSort(arr *[]int) {
	var heap Heap
	n := len(*arr)
	for i := 0; i < n; i++ {
		heap.Add((*arr)[i])
	}
	for i := 0; i < n; i++ {
		(*arr)[i] = heap.Top()
		if heap.Size() > 1 {
			heap.Set(0, heap.Remove())
		}

		heap.siftDown()
	}
}

func MergeSort(arr *[]int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	MergeSort(arr, left, mid)
	MergeSort(arr, mid + 1, right)
	lArr := make([]int, mid - left + 1)
	rArr := make([]int, right - mid)
	copy(lArr, (*arr)[left: mid + 1])
	copy(rArr, (*arr)[mid + 1: right + 1])
	k := left
	i := 0
	j := 0
	for i < mid - left + 1 && j < right - mid{
		if lArr[i] < rArr[j] {
			(*arr)[k] = lArr[i]
			k++
			i++
		} else {
			(*arr)[k] = rArr[j]
			k++
			j++
		}
	}
	for i < mid - left + 1 {
		(*arr)[k] = lArr[i]
		k++
		i++
	}
	for j < right - mid {
		(*arr)[k] = rArr[j]
		k++
		j++
	}
}

func BubbleSort(arr *[]int) {
	n := len(*arr)
	isSorted := true
	for i := 0; i < n; i++ {
		isSorted = true
		for j := i; j < n - 1; j++ {
			if (*arr)[j] > (*arr)[j + 1] {
				(*arr)[j], (*arr)[j + 1] = (*arr)[j + 1], (*arr)[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
}

func SelectSort(arr *[]int) {
	n := len(*arr)
	var min int
	for i := 0; i < n; i++ {
		min = i
		for j := i + 1; j < n; j++ {
			if (*arr)[j] < (*arr)[min] {
				min = j
			}
		}
		(*arr)[min], (*arr)[i] = (*arr)[i], (*arr)[min]
	}
}

func InsertSort(arr *[]int) {
	n := len(*arr)
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if (*arr)[j] <= (*arr)[j + 1] {
				break
			}
			(*arr)[j], (*arr)[j + 1] = (*arr)[j + 1], (*arr)[j]
		}
	}
}

func ShellSort(arr *[]int) {
	n := len(*arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i += gap {
			for j := i - gap; j >= 0; j-- {
				if (*arr)[j] <= (*arr)[j + gap] {
					break
				}
				(*arr)[j], (*arr)[j + gap]  =(*arr)[j + gap], (*arr)[j]
			}
		}
	}
}

func RadixSort(arr *[]int) {
	var bucket []Deque = make([]Deque, 10)
	for i := 0; i < 10; i++ {
		bucket[i] = Deque{Element: make([]int, 0), Len: 0}
	}
	var num int
	n := len(*arr)
	mod := 10
	isChanged := false
	for {
		isChanged = false
		for i := 0; i < n; i++ {
			num = (*arr)[i] % mod
			for num < 0 {
				num += 10
			}
			num /= mod / 10
			if num != 0 {
				isChanged = true
			}
			bucket[num].Add((*arr)[i])
		}
		k := 0
		for i := 0; i < 10; i++ {
			for !bucket[i].IsEmpty() {
				(*arr)[k] = bucket[i].Remove()
				k++
			}
			
		}
		
		if !isChanged {
			break
		}
		mod *= 10
	}
}

