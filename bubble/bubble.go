package main

import (
	"fmt"
)

// 若判断出数列已经有序，并做出标记，那么剩下的几轮排序就不必执行了，可以提前结束工作。
// 每一轮排序后，记录下来最后一次元素交换的位置，该位置即为无序数列的边界，
// 再往后就是有序区了。有序区则没有必要继续往下执行了
func bubbleSort(sli []int)[]int{
	doSwap := true
	lastUnsortedIndex := len(sli) - 1
	swapIndex := -1
	for doSwap {
		doSwap = false;
		for i := 0; i < lastUnsortedIndex; i++ {
			if sli[i] > sli[i+1] {
				sli[i], sli[i+1] = sli[i+1], sli[i]
			}
			doSwap = true
			swapIndex = i
		}
		lastUnsortedIndex = swapIndex
	}
	return sli
}

func main(){
	input := []int{9,6,7,2,1,8}
	fmt.Println("input:", input)
	output := bubbleSort(input)
	fmt.Println("output:", output)
}