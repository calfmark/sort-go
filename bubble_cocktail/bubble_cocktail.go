package gsort

// 鸡尾酒排序
// 排序过程就像钟摆一样，第1轮从左到右，第 2 轮从右到左，第 3 轮再从左到右……
// 目的就是解决像无序数列 {2,3,4,5,6,7,8,1}这样，只有少数元素位置不对的情况
// 实现类似冒泡算法：
// 若判断出数列已经有序，并做出标记，那么剩下的几轮排序就不必执行了，可以提前结束工作。
// 每一轮排序后，记录下来最后一次元素交换的位置，该位置即为无序数列的边界，
// 再往后就是有序区了。有序区则没有必要继续往下执行了
func Sort(sli []int) []int {
	doSwap := true

	//右无序数列的边界，每次比较只需要比到这里为止
	rightLastUnsortedIndex := len(sli) - 1
	//左无序数列的边界，每次比较只需要比到这里为止
	leftLastUnsortedIndex := 0
	for doSwap { //若没有发生交换，则认为已经有序
		swapIndex := -1
		doSwap = false
		for i := leftLastUnsortedIndex; i < rightLastUnsortedIndex; i++ {
			if sli[i] > sli[i+1] {
				sli[i], sli[i+1] = sli[i+1], sli[i]
				doSwap = true
				swapIndex = i
			}
		}
		rightLastUnsortedIndex = swapIndex
		if !doSwap { //若没有发生交换，则认为已经有序
			break
		}

		doSwap = false
		for j := rightLastUnsortedIndex; j > leftLastUnsortedIndex; j-- {
			if sli[j] < sli[j-1] {
				sli[j], sli[j-1] = sli[j-1], sli[j]
			}
			doSwap = true
			swapIndex = j
		}
		leftLastUnsortedIndex = swapIndex
	}
	return sli
}
/*
func main() {
	input := []int{9, 6, 7, 2, 1, 8}
	fmt.Println("input:", input)
	output := Sort(input)
	fmt.Println("output:", output)
}
*/
