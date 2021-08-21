package gsort

// 计数排序
// 基本思想：
// 第一步, 根据数组下标统计无序数列
// 第二步，从统计数组的第 2 个元素开始，每一个元素都加上前面所有元素之和。
//        这样相加的目的，是让统计数组存储的元素值，等于相应整数的最终排序位置的最后序号
// 第三步，从后向前遍历输入无序数列，根据其值，找到计数数组对应的值，也就是其最终的排序序号，
//        然后将计数数组对应元素值减一，表示下次若再次遇到相同的值，其排名就是当前的值，
//     	  比之前的排名靠前一位，保证了数列是稳定的排序
// 注意，此排序支持正负整数排序，不支持小数排序，若需要支持小数，需要先处理小数为整数进行排序，最后恢复为小数
func Sort(sli []int) []int {
	if len(sli) < 2 {
		return sli
	}
	// 1. 找到最大最小值
	max, min := sli[0], sli[0]
	for _, v := range sli {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	//2. 确定计数范围 ，下标 0~range-1 对应数字 min~max
	rang := max - min + 1
	countSli := make([]int, rang)

	//3. 将每个整数出现的次数统计到计数数组中对应下标的位置，
	//   这里需要将每个元素减去 min，才能映射到 0～range-1 范围内
	for _, v := range sli {
		countSli[v-min]++
	}

	//4. 每个元素在结果数组中的最后一个下标位置 = 前面比自己小的数字的总数 + 自己的数量 - 1。
	//   我们将 countSli[0] 先减去 1，后续 countSli 直接累加即可
	countSli[0]--
	for i := 1; i < rang; i++ {
		// 将 countSli 计算成当前数字在结果中的最后一个下标位置。
		// 位置 = 前面比自己小的数字的总数 + 自己的数量 - 1
		countSli[i] += countSli[i-1]
	}

	//5. 从后往前遍历数组，通过 countSli 中记录的下标位置，将 sli 中的元素放到 tmp数组中
	tmpSli := make([]int, len(sli))
	for i := len(sli) - 1; i >= 0; i-- {
		// countSli[sli[i] - min] 表示此元素在结果数组中的下标,
		// 由于前面的countSli[0] 已经减了 1，所以这里的减 1 可以省略
		tmpSli[countSli[sli[i]-min]] = sli[i]
		// 更新 countSli[sli[i] - min]，指向此元素的前一个下标
		countSli[sli[i]-min]--
	}

	copy(sli, tmpSli)
	return sli
}
