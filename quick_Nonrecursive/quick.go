package gsort

import (
	"math/rand"
	"time"
)

type node struct {
	start int
	end   int
}

// 快速排序-非递归实现
// 基本思想：
// 		从数组中取出一个数，称之为基数（pivot）
//		遍历数组，将比基数大的数字放到它的右边，比基数小的数字放到它的左边。遍历完成后，数组被分成了左右两个区域
//		将左右两个区域视为两个数组，重复前两个步骤，直到排序完成
func Sort(sli []int) []int {
	length := len(sli)
	if length <= 1 {
		return sli
	}
	stack := []node{}
	stack = append(stack, node{start: 0, end: length - 1})

	for len(stack) != 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		mid := partition(sli, item.start, item.end)
		//待排序的分区内至少要有2个元素
		if mid-1 > item.start {
			stack = append(stack, node{start: item.start, end: mid - 1})
		}
		if mid+1 < item.end {
			stack = append(stack, node{start: mid + 1, end: item.end})
		}
	}
	return sli
}

func partition(sli []int, start int, end int) int {
	//随机取一个随机元素作为基准元素，优化有序的序列情况
	//TODO 取随机数耗时可能过长，待优化
	randIndex := randIdx(start, end)
	if randIndex != start {
		sli[start], sli[randIndex] = sli[randIndex], sli[start]
	}

	pivot := sli[start]
	//mark表示小于基准元素的区域边界
	mark := start

	for i := start + 1; i <= end; i++ {
		if sli[i] < pivot {
			//小于基准元素的边界扩大1
			mark++
			sli[mark], sli[i] = sli[i], sli[mark]
		}
	}
	sli[start] = sli[mark]
	sli[mark] = pivot
	return mark
}

func randIdx(start int, end int) int {
	if start >= end {
		return start
	}
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	return rng.Intn(end-start) + start
}
