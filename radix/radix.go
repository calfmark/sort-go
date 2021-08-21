package gsort

import "math"

// 基数排序
// 基本思想：
// 1.找出数组中最大的数字的位数 maxDigitLength
// 2.获取数组中每个数字的基数
// 3.遍历 maxDigitLength 轮数组，每轮按照基数对其进行排序
// 注意，此排序支持正负整数排序，不支持小数排序，若需要支持小数，需要先处理小数为整数进行排序，最后恢复为小数
func Sort(sli []int) []int {
	if len(sli) < 2 {
		return sli
	}
	// 1. 找到最大绝对值
	max := int(math.Abs(float64(sli[0])))
	for _, v := range sli {
		vabs := int(math.Abs(float64(v)))
		if vabs > max {
			max = vabs
		}
	}

	//2. maxDigitLength
	maxDigitLength := 0
	for ; max != 0; max /= 10 {
		maxDigitLength++
	}
	tmpSli := make([]int, len(sli))
	dev := 1
	for ; maxDigitLength > 0; maxDigitLength-- {
		// 计数排序
		// 使用计数排序算法对基数进行排序，下标 [0, 18] 对应基数 [-9, 9]
		counting := make([]int, 19)
		for _, v := range sli {
			idx := v/dev%10 + 9
			counting[idx]++
		}
		// 每个元素在结果数组中的最后一个下标位置 = 前面比自己小的数字的总数 + 自己的数量 - 1。
		//  我们将 counting[0] 先减去 1，后续 counting 直接累加即可
		counting[0]--
		for i := 1; i < len(counting); i++ {
			// 将 counting 计算成当前数字在结果中的最后一个下标位置。
			// 位置 = 前面比自己小的数字的总数 + 自己的数量 - 1
			counting[i] += counting[i-1]
		}
		// 从后往前遍历数组，通过 counting 中记录的下标位置，将 sli中的元素放到 tmp数组中
		for i := len(sli) - 1; i >= 0; i-- {
			// counting[sli[i]/dev%10 + 9] 表示此元素在结果数组中的下标,
			idx := sli[i]/dev%10 + 9
			// 由于前面的counting[0] 已经减了 1，所以这里的减 1 可以省略
			tmpSli[counting[idx]] = sli[i]
			// 更新 counting[arr[i] - min]，指向此元素的前一个下标
			counting[idx]--
		}
		copy(sli, tmpSli)
		dev *= 10
	}

	return sli
}
