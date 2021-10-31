// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
    m := map[int]int{ nums[0]: 0 }

    for i := 1; i < len(nums); i++ {
        diff := target - nums[i]
        if j, ok := m[diff]; ok {
            return []int{i, j}
        }
        m[nums[i]] = i
    }
    return nil
}
