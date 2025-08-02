package main

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, num := range nums {
		val := target - num
		if j, ok := hash[val]; ok {
			return []int{j, i}
		}
		hash[num] = i
	}

	return []int{}
}
