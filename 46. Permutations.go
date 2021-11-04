// https://leetcode.com/problems/permutations/
func permute(nums []int) [][]int {
    var res [][]int
    helper(nums, []int{}, []int{}, &res)
    return res
}

func helper(nums []int, result []int, idxes []int, results *[][]int) {
    if len(result) == len(nums) {
        *results = append(*results, result)
        return
    }
    
    for i, n := range nums {
        if inIdxes(idxes, i) {
            continue
        }
        
        helper(nums, append(result, n), append(idxes, i), results)
    }
}

func inIdxes(idxes []int, i int) bool {
    for _, idx := range idxes {
        if i == idx {
            return true
        }
    }
    return false
}
