package main

/*
  缺失数字
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
*/

func missingNumber(nums []int) int {
	sum := len(nums) * (len(nums) + 1) / 2
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}
	return sum
	// xor := 0
	// for i := 0; i < len(nums); i++ {
	// 	xor ^= nums[i] ^ (i + 1)
	// }
	// return xor
	// sort.Ints(nums)
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] != i {
	// 		return i
	// 	}
	// }
	// return len(nums)
	// cache := make(map[int]struct{}, len(nums))
	// for i := 0; i < len(nums); i++ {
	// 	cache[nums[i]] = struct{}{}
	// }
	// for i := 0; i < len(nums); i++ {
	// 	_, exist := cache[i]
	// 	if !exist {
	// 		return i
	// 	}
	// }
	// return len(nums)
}

func main() {

}
