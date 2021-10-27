// https://leetcode.com/problems/climbing-stairs/
func climbStairs(n int) int {
    ans := 1
    lastAns := 1
    for i := 0; i < n; i++ {
        lastAns = ans + lastAns
        ans = lastAns - ans
    }
    return ans
}
