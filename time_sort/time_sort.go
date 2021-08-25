package gsort

// timeSort排序
// TimSort 是 Tim Peters 在 2002年提出的一种算法，它是归并算法的一种改进算法
// 基本思想：
// TimSort 的主要思想是：通过遍历数组，将数组拆分成若干个单调递增的子数组。
// 每一块称为一个 run。拆分完成后，再将 run 两两合并起来。
// 在遍历数组时，如果遇到单调递减的小块，TimSort 会将其翻转使其单调递增。
// TimSort 在对部分有序的数组进行排序时，速度很快。而现实世界中的数据往往总是部分有序的
// 参考：https://leetcode-cn.com/leetbook/read/sort-algorithms/phga33/
func Sort(sli []int) []int {
	left, right := 0, len(sli)-1
	if len(sli) < 2 {
		return sli
	}

	// 1. 将数组拆分为单调递增的小块
	run := make([]int, 1)
	run[0] = left
	count := 0
	for k := left; k < right; k++ {
		if sli[k] <= sli[k+1] { // 递增
			for k++; k < right && sli[k] <= sli[k+1]; k++ {
			}
		} else { // 递减
			for k++; k < right && sli[k] > sli[k+1]; k++ {
			}
			for lo, hi := run[count], k; lo < hi; {
				sli[lo], sli[hi] = sli[hi], sli[lo]
				hi--
				lo++
			}
		}
		count++
		run = append(run, k+1) //下一个小块的起始元素
	}
	if run[count] == right { //最后一个小块只包含一个元素
		count++
		run = append(run, right+1)
	} else if count == 1 { // 数组已经有序
		return sli
	}

	// 2. 相连两个小块，两两合并, 一个小块的区间是[run[k], run[k+1])
	tmpSli := make([]int, len(sli))
	aSli := sli
	bSli := tmpSli
	swapNum := 0
	for last := 0; count > 1; count = last {
		last = 0
		for k := last + 2; k <= count; k += 2 {
			mid, hi := run[k-1], run[k]
			i := run[k-2]
			for p0, p1 := i, mid; i < hi; i++ {
				if p1 >= hi || p0 < mid && aSli[p0] <= aSli[p1] {
					bSli[i] = aSli[p0]
					p0++
				} else {
					bSli[i] = aSli[p1]
					p1++
				}
			}
			last++
			run[last] = hi //相连两个小块合并为一个小块
		}
		if count&1 != 0 { //块数为奇数时,最后一个小块则不用合并, 注意count为run的最大下标也是小块的个数
			for i, lo := right, run[count-1]; i >= lo; i-- {
				bSli[i] = aSli[i]
			}
			last++
			run[last] = right + 1
		}
		aSli, bSli = bSli, aSli
		swapNum++
	}
	// 3. 若交换了奇数次，排完序的元素在tmpSli里
	if swapNum&1 != 0 {
		copy(sli, tmpSli)
	}

	return sli
}
