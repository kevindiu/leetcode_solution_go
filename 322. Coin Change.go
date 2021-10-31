// https://leetcode.com/problems/coin-change
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount + 1)
    for i := 1; i < amount + 1; i++ {
        dp[i] = amount + 1
    }
    dp[0] = 0
    
    // sort the coins for the shortcut
    sort.Ints(coins)
    
    for i := 1; i <= amount; i++ {
        for _, c := range coins {
            diff := i-c
            // since it is sorted, we only need to calcuate until i-c is negative value
            if diff < 0 {
                break
            }
            dp[i] = min(dp[i], 1 + dp[i-c])
        }
    }
    
    if dp[amount] >= amount + 1 {
        return -1
    }
    return dp[amount]
}

func min(n1, n2 int) int {
    if n1 > n2 {
        return n2
    }
    return n1
}
