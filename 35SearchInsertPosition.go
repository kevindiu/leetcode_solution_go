// https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
    le := len(nums)
    if le == 0 {
        return 0
    }
        
    start := 0
    end := le - 1
    for start <= end {
        mid := start + ((end - start) / 2)
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target { 
            start = mid + 1
        } else {
            end = mid - 1
        }
    }
    
    return start
}
