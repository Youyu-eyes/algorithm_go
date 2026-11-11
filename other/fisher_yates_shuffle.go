package other

import (
	"math/rand/v2"
)

func FisherYatesShuffle(nums []int) []int {
	for i := range nums {
		j := rand.IntN(i)
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}