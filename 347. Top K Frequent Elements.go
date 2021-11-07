// https://leetcode.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
    // count the frequency of each nums
    freqMap := make(map[int]int)
    for _, n := range nums {
        freqMap[n]++
    }
    
    // convert the frequncy map to bucket
    buckets := make([][]int, len(nums) + 1)
    for n, f := range freqMap {
        buckets[f] = append(buckets[f], n)
    }
    
    // find the result from frequency map
    res := make([]int, 0, k)
    for i := len(buckets) - 1; i >= 0; i-- {
        if buckets[i] != nil {
            res = append(res, buckets[i]...)
        }
        if len(res) >= k {
            return res
        }
    }
    
    return res
}
