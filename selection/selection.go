package gsort

// 二元选择排序
//每轮选择时记录最小值和最大值，然后分别把最大最小值移到数组的头尾。
func Sort(sli []int) []int {
	len := len(sli)
	// i 只需要遍历一半
	for i := 0; i < len/2; i++ {
		maxIdx := i
		minIdx := i
		for j := i + 1; j < len-i; j++ {
			if sli[j] > sli[maxIdx] {
				maxIdx = j
			}
			if sli[j] < sli[minIdx] {
				minIdx = j
			}
		}
		// 如果 minIdx 和 maxIdx 都相等，那么他们必定都等于 i，
		// 且后面的所有数字都与 sli[i] 相等，此时已经排序完成
		if maxIdx == minIdx {
			break
		}
		// 将最小元素交换至首位
		if minIdx != i {
			sli[i], sli[minIdx] = sli[minIdx], sli[i]
		}

		// 如果最大值的下标刚好是 i，由于 sli[i] 和 sli[minIdx] 已经交换了，
		// 所以这里要更新 maxIdx 的值。
		if maxIdx == i {
			maxIdx = minIdx
		}
		// 将最大元素交换至末尾
		sli[maxIdx], sli[len-1-i] = sli[len-1-i], sli[maxIdx]
	}
	return sli
}
