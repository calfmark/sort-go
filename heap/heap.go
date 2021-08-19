package gsort

// 堆排序
// 把无序数组构建成二叉堆。需要从小到大排序，则构建成最大堆；需要从大到小排序，则构建成最小堆。
// 循环删除堆顶元素，替换到二叉堆的末尾，调整堆产生新的堆顶。
// 堆有如下性质：
// 		对于完全二叉树中的第 i 个数，它的左子节点下标：left = 2i + 1
//		对于完全二叉树中的第 i 个数，它的右子节点下标：right = left + 1 即right = 2i + 2
//		对于有 n 个元素的完全二叉树(n≥2)(n≥2)，它的最后一个非叶子结点的下标：n/2 - 1
func Sort(sli []int) []int {
	slen := len(sli)
	// 1.把无序数组构建成最大堆
	buildHeap(sli)
	 // 2.循环删除堆顶元素，移到集合尾部，调整堆产生新的堆顶
	for i:=slen-1; i > 0; i-- {
		// 堆顶交换到数组末尾, 即最后1个元素和第1个元素进行交换
		sli[0], sli[i] = sli[i], sli[0]
		// 调整二叉树为最大堆
		downAdjust(sli, 0, i)
	}

	return sli
}

func buildHeap(sli []int){
	slen := len(sli)
	for i:= (slen>>1)-1; i >= 0; i-- {
		downAdjust(sli, i, slen)
	}

}
//下沉
func downAdjust(sli []int, parentIdx int, len int){
	childIdx := (parentIdx<<1) + 1
	// 保存父节点值，用于最后的赋值
	tmp := sli[parentIdx]
	for childIdx < len {
		// 如果有右孩子，且右孩子大于左孩子的值，则定位到右孩子
		if (childIdx + 1) < len && sli[childIdx+1] > sli[childIdx]{
			childIdx++
		}
		// 如果父节点大于任何一个孩子的值，则直接跳出
		if tmp >= sli[childIdx] {
			break
		}
		//无须真正交换，单向赋值即可
		sli[parentIdx] = sli[childIdx]
		parentIdx = childIdx
		childIdx = (parentIdx << 1) + 1
	}
	// 将tmp值保存到最后对应腾出的位置
	sli[parentIdx] = tmp
}
