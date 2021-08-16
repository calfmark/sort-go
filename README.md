排序算法
====
本项目是个人的golang学习项目，最近学了golang，记录一下练习题。整理出用golang实现的排序算法。主要是参考了以下教程：
- [leetcode-排序算法全解析](https://leetcode-cn.com/leetbook/read/sort-algorithms/evdcgv/)
- [漫画算法-小灰的算法之旅](https://leetcode-cn.com/leetbook/read/journey-of-algorithm/5eay2g/)

代码环境使用的环境是：`go version go1.16.2 windows/amd64`

**注意:** 某个排序算法有很多种实现，但这里只会实现最优的版

### 测试方法

1. 单个排序算法测试方法，直接进入到某个排序目录执行：`go test -v` 或者 `go test -v -bench=. -benchmem`
2. 一键整体测试，在根目录下执行`sh test.sh`