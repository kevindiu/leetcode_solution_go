// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
    numLen := len(nums)
    if numLen == 0 {
        return []int {-1, -1}
    }
    
    start := findFirstIdx(nums, target)
    if start == -1 {
        return []int {-1, -1}
    }
    end := findEndIdx(nums, start, target)
   
    return []int {start, end}
}

func findFirstIdx(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    idx := -1
    for left <= right {
        mid := left + ((right - left) / 2)
        if nums[mid] == target {
            idx = mid
        } 
        
        if nums[mid] >= target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return idx
}

func findEndIdx(nums []int, startIdx, target int) int {
    left, right := startIdx, len(nums) - 1
    idx := startIdx
    for left <= right {
        mid := left + ((right - left) / 2)
        if nums[mid] == target {
            idx = mid
        } 
        
        if nums[mid] <= target {
            left = mid +1
        } else {
            right = mid -1
        }
    }
    return idx
}
