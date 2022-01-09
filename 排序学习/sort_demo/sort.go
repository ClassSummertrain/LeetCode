package main

import (
	"fmt"
	"math"
)

//bublleSort:
//思想：依次比较相邻记录大小，如果是逆序就交换，直到没有逆序
//代码1：未剪枝，妥妥的n^2
func bubbleSort1(array []int) {
	for i := 0; i < len(array)-1; i++ {
		//每一轮会将最大的元素冒泡到最后有序区的合适位置
		for j := 0; j < len(array)-2; j++ {
			if array[j] > array[j+1] {
				tmp := array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
			}
		}
	}
}

//代码1：有一定剪枝，在i,j处可减少一定次数
func bubbleSort2(array []int) {
	//每一轮会将最大的元素冒泡到最后有序区的合适位置
	for i := len(array) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				tmp := array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
			}
		}
	}
}

//insertSort:
//思想：初始第一个为有序区，每轮将后面无序区第一个元素插入有序区
func insertSort(array []int) {
	for i := 1; i < len(array)-1; i++ {
		key := array[i]
		j := i - 1
		//insert key to sorted sequence array[0...i-1]
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		//注意是j+1
		array[j+1] = key
	}
}

//mergeSort:
//思想：将待排序的序列划分为两个等长子序列，对这两个序列排序
//得到两个有序子序列，再将这两个子序列合并成一个有序序列

//merge用来进行两个有序子序列的合并
//p,q,r   : array[p...q],array[q+1,r]
func merge(array []int, p, q, r int) {
	//左右字序列长度
	len1, len2 := q-p+1, r-q
	//为左右子序列建立新数组，避免原地修改破坏原数组
	left := make([]int, len1)
	right := make([]int, len2)
	//copy
	copy(left[:], array[p:q+1])
	copy(right[:], array[q+1:r+1])
	//为左右子序列设置哨兵，替代越界检查操作
	left, right = append(left, math.MaxInt), append(right, math.MaxInt)
	i, j := 0, 0
	//左右子序列合并成有序序列
	for k := p; k <= r; k++ {
		if left[i] <= right[j] {
			array[k] = left[i]
			i++
		} else {
			array[k] = right[j]
			j++
		}
	}
}

//排序子序列array[p...r]中的元素
func mergeSort(array []int, p, r int) {
	if p < r {
		//递归划分左右子序列
		q := (p + r) / 2
		mergeSort(array, p, q)
		mergeSort(array, q+1, r)
		//归并
		merge(array, p, q, r)
	}
}

//heapSort:
//思想：调整堆，建堆。
//modify思想：假定左右子树均已经是最大堆，
//通过modify(A[i])在堆中逐渐下降以遵循最大堆性质
func modify(array []int, i, heapSize int) {
	//左右孩子位置，交换函数
	left := func() int { return 2 * i }
	right := func() int { return 2*i + 1 }
	exchang := func(x, y *int) {
		tmp := *x
		*x = *y
		*y = tmp
	}
	//堆大小 临时变量(最大值的下标)
	largest := 0
	//找到根，左,右中最大值的下标
	if left() < heapSize && array[left()] > array[i] {
		largest = left()
	} else {
		largest = i
	}
	if right() < heapSize && array[right()] > array[largest] {
		largest = right()
	}
	//调整
	if largest != i {
		exchang(&array[i], &array[largest])
		modify(array, largest, heapSize)
	}
}

func heapSort(array []int) {
	heapSize := len(array)
	exchang := func(x, y *int) {
		tmp := *x
		*x = *y
		*y = tmp
	}
	buildMaxHeap := func(array []int) {
		for i := len(array) / 2; i >= 0; i-- {
			modify(array, i, heapSize)
		}
	}
	buildMaxHeap(array)
	for i := len(array) - 1; i >= 0; i-- {
		exchang(&array[0], &array[i])
		heapSize -= 1
		modify(array, 0, heapSize)
	}
}

//quickSort:
//思想：快排划分子序列，对子序列继续调用快排划分
func quickSort(array []int, p, r int) {
	exchang := func(x, y *int) {
		tmp := *x
		*x = *y
		*y = tmp
	}
	partition := func(array []int, p, r int) int {
		key := array[r]
		i := p - 1
		for j := p; j < r; j++ {
			if array[j] <= key {
				i++
				exchang(&array[i], &array[j])
			}
		}
		exchang(&array[i+1], &array[r])
		return i + 1
	}
	if p < r {
		q := partition(array[:], p, r)
		quickSort(array[:], p, q-1)
		quickSort(array[:], q+1, r)
	}
}

//countingSort:
//思想：对输入数组的每一个数，确定比他小的元素个数，
//根据这个结果，在输出数组相应位置直接进行插入
//k为top border ,include k
func countingSort(array []int, k int) {
	//assitant array ,innitialize:C[i]=0
	C := make([]int, k+1)
	//C[i] is the number of array[i]=i
	for i := 0; i < len(array); i++ {
		C[array[i]] += 1
	}
	//C[i] is the number of  range array[i]<=i
	for i := 1; i < len(C); i++ {
		C[i] += C[i-1]
	}
	//B is the result array,innitialize:B[i]=0
	B := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		B[C[array[i]]-1] = array[i]
		C[array[i]] -= 1
	}
	copy(array, B)
}

//radixSort：
//思想：每次对一个关键码排序，（主关键码优先，次关键码优先）
func radixSort(array []int) {
	length := len(array)
	if length <= 1 {
		return
	}
	//最大数
	max := array[0]
	for _, val := range array {
		if val > max {
			max = val
		}
	}
	//最大位
	digit := 0
	for max > 0 {
		max /= 10
		digit++
	}

	//桶
	buket := [10][]int{}

	index := 1
	for i := 1; i <= digit; i++ {
		for _, val := range array {
			tmp := (val / index) % 10
			buket[tmp] = append(buket[tmp], val)
		}
		newIndex := 0
		for b, ints := range buket {
			for _, val := range ints {
				array[newIndex] = val
				newIndex++
			}
			//清空桶buket[b]
			buket[b] = []int{}
		}
		index *= 10
	}
}

//buketSort:
//思想：分类收集，和基数排序差不多，差别我还在思考
// func buketSort(array []int) {
// }
func main() {
	var array1 [6]int = [6]int{11, 25, 33, 32, 24, 96}
	bubbleSort1(array1[:])
	fmt.Println(array1)
	var array2 [6]int = [6]int{11, 25, 33, 32, 24, 96}
	bubbleSort2(array2[:])
	fmt.Println(array2)
	var array3 [6]int = [6]int{11, 25, 33, 32, 24, 96}
	insertSort(array3[:])
	fmt.Println(array3)
	var array4 [6]int = [6]int{11, 25, 33, 32, 24, 96}
	mergeSort(array4[:], 0, len(array4)-1)
	fmt.Println(array4)
	var array5 [6]int = [6]int{11, 25, 33, 32, 24, 96}
	heapSort(array5[:])
	fmt.Println(array5)
	var array6 [6]int = [6]int{11, 25, 33, 32, 24, 96}
	quickSort(array6[:], 0, len(array6)-1)
	fmt.Println(array6)
	var array7 [6]int = [6]int{11, 25, 33, 32, 24, 96}
	countingSort(array7[:], 96)
	fmt.Println(array7)
	var array8 [6]int = [6]int{11, 25, 33, 32, 24, 96}
	radixSort(array8[:])
	fmt.Println(array8)
	// var array9 [6]int = [6]int{11, 25, 33, 32, 24, 96}
	// buketSort(array9[:])
	// //fmt.Println(array8)

}
