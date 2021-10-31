// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
    l := len(nums)
    if l == 0 {
        return 0
    }
    
    i := 0
    for j := 1; j < l; j++ {
        if nums[j] != nums[i] {
            i ++
            nums[i] = nums[j]
        }
    }
    
    return i + 1
}

