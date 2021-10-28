// https://leetcode.com/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int)  {
    i1, i2 := m-1, n-1
    j := m + n -1
    for i1 >= 0 && i2 >= 0 {
        n1 := nums1[i1]
        n2 := nums2[i2]
        if n2 > n1 {
            nums1[j] = n2
            i2--
        } else {
            nums1[j] = n1
            i1-- 
        }
        j--
    }
    // process non-copied values in nums2
    for i2 >= 0 {
        nums1[j] = nums2[i2]
        i2 --
        j --
    }
    
}
