// https://leetcode.com/problems/plus-one/
func plusOne(digits []int) []int {
    for i := len(digits)-1; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i]++
            break
        }    
    
        if i == 0 {
            res := make([]int, len(digits) + 1)
            res[0] = 1
            return res   
        }
        digits[i] = 0
    }
    return digits
}
