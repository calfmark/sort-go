package gsort

// 归并排序-非递归实现
// 基本思想：
// 归并排序分成两步，一是拆分数组，二是合并数组
// 即将 1 个数字组成的有序数组合并成一个包含 2 个数字的有序数组，
// 再将 2 个数字组成的有序数组合并成包含 4 个数字的有序数组...直到整个数组排序完成
// 非递归则是先只包含1个数字组成的数组开始合并，合并到整个数组长度的有序数组
func Sort(sli []int) []int {
	if len(sli) <= 1 {
		return sli
	}
	tmpSli := make([]int, len(sli))
	mergeSort(sli, 0, len(sli)-1, tmpSli)
	return sli
}

func mergeSort(sli []int, start int, end int, tmpSli []int) {
	//有序子数组的长度
	length := 1

	for length < len(sli) {
		mergePass(sli, length, tmpSli)
		length *= 2
	}
}

// 两两归并相连有序子数组
// length 为有序子数组的长度
func mergePass(sli []int, length int, tmpSli []int) {
	i := 0
	for ; i <= len(sli)-2*length; i += 2 * length {
		merge(sli, i, i+2*length-1, i+length-1, tmpSli)
	}
	if (i + length) < len(sli) {
		// i+一个子数组的长度还小于数组长度，说明有两个长度不等的子数组
		// 其中一个长度为length，另外一个小于length
		merge(sli, i, len(sli)-1, i+length-1, tmpSli) //归并最后两个子数组
	}
	//不足一个length长度的子数组无需处理

}

// 将两个有序数组合并成一个
// 注意：这里的mid必须由上层传入，不能自动生成，这是跟递归实现的主要差别
func merge(sli []int, start int, end int, mid int, tmpSli []int) {
	//mid := (start + end) >> 1
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
