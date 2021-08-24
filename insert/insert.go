package gsort

// 插入排序
// 移动法：在新数字插入过程中，与前面的数字不断比较，前面的数字不断向后挪出位置，
// 当新数字找到自己的位置后，插入一次即可。
// 整个过程就像是已经有一些数字坐成了一排，这时一个新的数字要加入，
// 所以这一排数字不断地向后腾出位置，当新的数字找到自己合适的位置后，
// 就可以直接坐下了。重复此过程，直到排序结束。
func Sort(sli []int) []int {
	len := len(sli)
	for i := 1; i < len; i++ {
		//待往前插入的数站起来
		currentV := sli[i]
		j := i
		for ; j > 0 && currentV < sli[j-1]; j-- {
			sli[j] = sli[j-1]
		}
		//找到位置坐下
		sli[j] = currentV
	}
	return sli
}
