package gsort

// 双插入排序
// 每次从左边取两个数字，大的数字记录为 big，小的数字记为 small，然后不断向前寻找 big 的插入位置，
// 当 big 插入后，small 直接从 big 的插入位置开始，向前寻找自己的插入位置。
// 最后一个数字可能需要单独插入。
// note: 这里实现非哨兵的方式
// 具体原理可以参考：https://leetcode-cn.com/leetbook/read/sort-algorithms/phzr8m/
// 根据java sort里的实现, 当数据量小于47个时，采用插入排序的时间复杂度应该是最优的
func Sort(sli []int) []int {
	if len(sli) < 2 {
		return sli
	}

	left, right := 1, len(sli)-1
	//忽略前面已有序的部分
	for ; left <= right && sli[left] >= sli[left-1]; left++ {
	}

	k := left
	for left++; left <= right; left += 2 {
		//将待插入的连续的两个数站起来
		big, small := sli[k], sli[k+1]
		if big < small {
			big, small = small, big
		}
		//找到较大数插入的位置
		for k--; k >= 0 && big < sli[k]; k-- {
			sli[k+2] = sli[k]
		}
		sli[k+2] = big
		//找到较小数插入的位置
		for ; k >= 0 && small < sli[k]; k-- {
			sli[k+1] = sli[k]
		}
		sli[k+1] = small
		k = left + 1
	}
	//处理最后一个元素
	if k == right {
		last := sli[right]
		for ; right > 0 && last < sli[right-1]; right-- {
			sli[right] = sli[right-1]
		}
		sli[right] = last
	}

	return sli
}
