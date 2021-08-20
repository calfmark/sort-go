package gsort

// 归并排序
// 基本思想：
// 归并排序分成两步，一是拆分数组，二是合并数组
// 即将 1 个数字组成的有序数组合并成一个包含 2 个数字的有序数组，
// 再将 2 个数字组成的有序数组合并成包含 4 个数字的有序数组...直到整个数组排序完成
func Sort(sli []int) []int {
	tmpSli := make([]int, len(sli))
	mergeSort(sli, 0, len(sli)-1, tmpSli)
	return sli
}

func mergeSort(sli []int, start int, end int, tmpSli []int) {
	if start >= end || start < 0 || end < 0 {
		return
	}
	mid := (start + end) >> 1
	mergeSort(sli, start, mid, tmpSli)
	mergeSort(sli, mid+1, end, tmpSli)

	merge(sli, start, end, tmpSli)
}

// 将两个有序数组合并成一个
func merge(sli []int, start int, end int, tmpSli []int) {
	mid := (start + end) >> 1
	//数组1的首尾位置
	start1, end1 := start, mid
	//数组2的首尾位置
	start2, end2 := mid+1, end
	//遍历两个数组的索引
	index1, index2 := start1, start2
	resultIdx := start
	for index1 <= end1 && index2 <= end2 {
		if sli[index1] <= sli[index2] {
			tmpSli[resultIdx] = sli[index1]
			index1++
		} else {
			tmpSli[resultIdx] = sli[index2]
			index2++
		}
		resultIdx++
	}
	// 将剩余数字补到结果数组之后
	for ; index1 <= end1; index1++ {
		tmpSli[resultIdx] = sli[index1]
		resultIdx++
	}
	for ; index2 <= end2; index2++ {
		tmpSli[resultIdx] = sli[index2]
		resultIdx++
	}
	// 将 tmpSli 操作区间的数字拷贝到 sli 中，以便下次比较
	for i := start; i <= end; i++ {
		sli[i] = tmpSli[i]
	}
}
