package gsort

import (
	"sort"
)

// 桶排序
// 基本思想：
// 1. 将区间划分为 n 个相同大小的子区间，每个子区间称为一个桶
// 2. 遍历数组，将每个数字装入桶中
// 3. 对每个桶内的数字单独排序，这里需要采用其他排序算法，如插入、归并、快排等
// 4. 最后按照顺序将所有桶内的数字合并起来

func Sort(sli []int) []int {
	if len(sli) < 2 {
		return sli
	}
	// 1. 找到最大最小值
	min, max := sli[0], sli[0]
	for _, v := range sli {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	// 2. 计算桶和桶之间的距离
	rang := max - min
	bucketNum := len(sli) //可以指定任何的桶数
	// 桶和桶之间的距离
	// gap必须为flow类型,而且这里必须减1(可以把bucket理解为分割线会更好理解一些)，
	// 最后一个桶只会存max一个元素
	var gap float32 = float32(float32(rang) / float32(bucketNum-1))

	//3. 装桶
	buckets := make([][]int, bucketNum)
	for _, v := range sli {
		idx := int(float32(v-min) / gap)
		/*
			if len(buckets[idx]) == 0 {
				buckets[idx] = make([]int, 0)
			}
		*/
		//nil的slice可以使用append
		buckets[idx] = append(buckets[idx], v)
	}

	//4.对每个桶单独排序并将结果进行合并
	index := 0
	for i := 0; i < bucketNum; i++ {
		if len(buckets[i]) > 0 {
			sort.Ints(buckets[i])
			copy(sli[index:], buckets[i])
			index += len(buckets[i])
		}
	}
	return sli
}
