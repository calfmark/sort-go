package gsort

// 双轴快排排序
// 基本思想：
// 双轴快排每轮选取两个轴 pivot1、pivot2 (pivot1 < pivot2)，然后以两个轴为分界，
// 将数组分为左中右三个区域。交换三个区域内的数字，使得三个区域分别属于区间
// (-∞, pivot1)、[pivot1, pivot2]、(pivot2, +∞)。然后再对左中右区域不断重复此过程，
// 直至排序完成。
// 具体步骤：
// 通过数组的中间位置往前和往后走两次数组长度的 1/7 步长取出 5个备选轴。
// 如果 5个备选轴互不相等，取第二个轴和第四个轴进行双轴快排，将数组分成三个区域：
//  (-∞, pivot1)、[pivot1, pivot2]、(pivot2, +∞)。
// 分区后，如果中间区域过大（大于数组长度的 4/7），则将中间区域再次分成三个区域：
//  [pivot1, pivo1]、(pivot1, pivot2)、[pivot2, pivot2]，
//  只让 (pivot1, pivot2)区间参与下一轮双轴快排。
// 最后，对左中右三个区域递归进行排序。
// 具体参见: https://leetcode-cn.com/leetbook/read/sort-algorithms/phs9u1/
func Sort(sli []int) []int {
	doubleQuickSort(sli, 0, (len(sli) - 1))
	return sli
}

func doubleQuickSort(sli []int, start int, end int) {
	length := end - start + 1
	if length < 2 {
		return
	}
	seventh := length / 7
	e3 := (start + end) >> 1
	e2 := e3 - seventh
	e1 := e2 - seventh
	e4 := e3 + seventh
	e5 := e4 + seventh
	insertSort(sli, e1, e2, e3, e4, e5)

	// (-∞, less)、[less, great]、(great, +∞)。
	//左边区域的后一个元素的下标
	less := start
	//右边区域的前一个元素的下标
	great := end
	if sli[e1] != sli[e2] && sli[e2] != sli[e3] && sli[e3] != sli[e4] && sli[e4] != sli[e5] {
		// 双轴快排
		pivot1 := sli[e2]
		pivot2 := sli[e4]
		// 将pivot1和pivot2放到首尾去，最后再将其与less-1和great+1交换
		sli[e2], sli[start] = sli[start], sli[e2]
		sli[e4], sli[end] = sli[end], sli[e4]

		// 忽略已经有序的部分
		for less++; sli[less] < pivot1; less++ {
		}
		for great--; sli[great] > pivot2; great-- {
		}
        /*
         * Partitioning:
         *
         *   left part           center part                   right part
         * +--------------------------------------------------------------+
         * |  < pivot1  |  pivot1 <= && <= pivot2  |    ?    |  > pivot2  |
         * +--------------------------------------------------------------+
         *               ^                          ^       ^
         *               |                          |       |
         *              less                        k     great
         *
         * Invariants:
         *
         *              all in (start, less)   < pivot1
         *    pivot1 <= all in [less, k)     <= pivot2
         *              all in (great, end) > pivot2
         *
         * Pointer k is the first index of ?-part.
         */
		outer:
		for k := less; k <= great; k++ {
			tmpk := sli[k]
			if tmpk < pivot1 { // 移到左边
				sli[k] = sli[less]
				sli[less] = tmpk
				less++
			} else if tmpk > pivot2 { // 移到右边
				// 因为sli[great]还没有判断，所以不知道它属于哪个区域
				for ; sli[great] > pivot2 && k >= great; great-- {
				}
				if great < k {
					break outer
				}
				// 此时great <= pivot2
				sli[k] = sli[great]
				sli[great] = tmpk
				great--
				// 这里将 k 减 1，在下一次循环时，k 会加 1，
				// 所以这里是指继续判断 k 位置的数字属于哪个区间
				// （这个位置的数字是刚交换过来的 a[great]）
				k--
			}
		}
		sli[start], sli[less-1] = sli[less-1], sli[start]
		sli[end], sli[great+1] = sli[great+1], sli[end]
		doubleQuickSort(sli, start, less-2)
		doubleQuickSort(sli, less, great)
		doubleQuickSort(sli, great+2, end)

	} else {
		// 单轴快排
		pivot := sli[e3]
		sli[start], sli[e3] = sli[e3], sli[start]
		//mark表示小于基准元素的区域边界
		mark := start

		for i := start + 1; i <= end; i++ {
			if sli[i] < pivot {
				//小于基准元素的边界扩大1
				mark++
				sli[mark], sli[i] = sli[i], sli[mark]
			}
		}
		sli[start] = sli[mark]
		sli[mark] = pivot

		doubleQuickSort(sli, start, mark-1)
		doubleQuickSort(sli, mark+1, end)
	}
}

// 使用插入排序对e1-e5排序
func insertSort(sli []int, e1, e2, e3, e4, e5 int) {
	if sli[e2] < sli[e1] {
		sli[e1], sli[e2] = sli[e2], sli[e1]
	}

	if sli[e3] < sli[e2] {
		sli[e2], sli[e3] = sli[e3], sli[e2]
		if sli[e2] < sli[e1] {
			sli[e1], sli[e2] = sli[e2], sli[e1]
		}
	}

	if sli[e4] < sli[e3] {
		sli[e3], sli[e4] = sli[e4], sli[e3]
		if sli[e3] < sli[e2] {
			sli[e2], sli[e3] = sli[e3], sli[e2]
			if sli[e2] < sli[e1] {
				sli[e1], sli[e2] = sli[e2], sli[e1]
			}
		}
	}

	if sli[e5] < sli[e4] {
		sli[e5], sli[e4] = sli[e4], sli[e5]
		if sli[e4] < sli[e3] {
			sli[e3], sli[e4] = sli[e4], sli[e3]
			if sli[e3] < sli[e2] {
				sli[e2], sli[e3] = sli[e3], sli[e2]
				if sli[e2] < sli[e1] {
					sli[e1], sli[e2] = sli[e2], sli[e1]
				}
			}
		}
	}
}
