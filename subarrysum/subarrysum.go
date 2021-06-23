package subarrysum

import "fmt"

//采用前缀和思路 时间复杂度O(n)
//公式推到 sum[i] = sum[i-1]+nums[i] 如果sum[i]-sum[j-1] = k sum[j-1] = sum[i] -k
func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	m := map[int]int{}//使用hash map是因为可以查找查找是否key
	m[0] = 1 //确保第一次出现 满足的时候 count能加一
	for i := 0; i < len(nums); i++ {
		sum += nums[i] //进行前缀和计算
		if _, ok := m[sum - k]; ok {
			count += m[sum - k]
		}
		m[sum]++ //前缀和记录在map中
	}
	return count
}

func Test() {
	a := []int{1,1,1,1,1}
	b:= subarraySum(a,2)
	fmt.Println(b)
}