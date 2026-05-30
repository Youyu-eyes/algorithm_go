package other

// leetcode 169.多数元素
// https://leetcode.cn/problems/majority-element/
func majorityElement(nums []int) int {
    hp, ans := 0, 0
    for _, x := range nums {
        if hp == 0 {
            ans, hp = x, 1
        } else if x == ans {
            hp++
        } else {
            hp--
        }
    }
    return ans
}