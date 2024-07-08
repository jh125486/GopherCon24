package main

func main() {
	// noop
}

func solution1(nums1 []int, nums2 []int) []int {
	everything := make(map[int]int)
	for i := range nums1 {
		everything[nums1[i]]++ //= nums1[i]
	}
	var answer []int
	for i := range nums2 {
		if everything[nums2[i]] > 0 {
			answer = append(answer, nums2[i])
			everything[nums2[i]]--
		}
	}
	return answer
}

func solution2(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	used := make([]bool, len(nums2))
	for _, num := range nums1 {
		for j, numj := range nums2 {
			if num == numj && !used[j] {
				result = append(result, num)
				used[j] = true
				break
			}
		}
	}
	return result
}

func solution3(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	counter1 := make(map[int]int, len(nums1))
	for _, n := range nums1 {
		counter1[n]++
	}
	counter2 := make(map[int]int, len(nums2))
	for _, n := range nums2 {
		counter2[n]++
	}
	intersection := make([]int, 0, len(nums1))
	for n, c := range counter1 {
		for i := 0; i < min(c, counter2[n]); i++ {
			intersection = append(intersection, n)
		}
	}
	return intersection
}

func solution4(nums1 []int, nums2 []int) []int {
	set := map[int]int{}
	for _, val := range nums1 {
		set[val] = set[val] + 1
	}
	k := 0
	for _, val := range nums2 {
		if set[val] > 0 {
			nums1[k] = val
			k++
			set[val]--
		}
	}
	return nums1[:k]
}
