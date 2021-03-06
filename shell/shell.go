package gsort

// 希尔排序
// 基本思想：
//  1. 将待排序数组按照一定的间隔分为多个子数组，每组分别进行插入排序。
//     这里按照间隔分组指的不是取连续的一段数组，而(必须)是每跳跃固定间隔gap
//     取一个值组成一组，也就是每个小组中的元素间隔固定为gap，也就分成了gap个小组，
//     换句话说就是将相距为gap的元素组成一一个子分组/序列,例如：
//     将一个长度为9的数组，分割为gap=9/2=4组的方法为：
//     第一组为{0,4,8},第二组为{1,5},第三组为{2,6},第四组为{3,7}，
//     也就是每个小组里的元素之间的间隔gap为4，也就是分成了4组
// 2. 逐渐缩小间隔进行下一轮排序
// 3. 最后一轮时，取间隔为 1，也就相当于直接使用插入排序。但这时经过前面的「宏观调控」，
//    数组已经基本有序了，所以此时的插入排序只需进行少量交换便可完成

//采用Knuth增量序列
func Sort(sli []int) []int {
	maxKnuth := 1
	len := len(sli)
	// 找到当前数组需要用到的 Knuth 序列中的最大值
	for maxKnuth <= len/3 {
		maxKnuth = 3*maxKnuth + 1
	}
	// gap增量按照 Knuth 序列规则依次递减到1为止
	for gap := maxKnuth; gap > 0; gap = (gap - 1) / 3 {
		// 采用移动插入排序对每个分组进行排序

		// 依次遍历每个组的第二个元素，也就是从第gap索引开始(即第二个元素开始)，
		// 按照顺序将每个元素依次向前插入自己所在的组
		for i := gap; i < len; i++ {
			// currentNumber 站起来，开始找位置
			currentV := sli[i]
			// 一个组中相连元素的下标索引间隔为gap
			j := i - gap
			for ; j >= 0 && currentV < sli[j]; j -= gap {
				// 向后挪位置
				sli[j+gap] = sli[j]
			}
			// currentV找到了自己的位置，坐下
			sli[j+gap] = currentV
		}
	}
	return sli
}
