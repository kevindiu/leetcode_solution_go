// https://leetcode.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
    max, sum := nums[0], 0
    for i := 0; i< len(nums); i++ {
        sum += nums[i]
        max = maxN(max, sum)
        if sum < 0 {
            sum = 0
        }
    }
    return max
}

func maxN(n1, n2 int) int {
    if n1 > n2 {
        return n1
    }
    return n2
}
