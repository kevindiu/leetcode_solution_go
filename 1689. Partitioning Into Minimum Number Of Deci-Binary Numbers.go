// https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
func minPartitions(n string) int {
    if len(n) == 0 {
        return 0
    }
    max := n[0]
    for i := 1; i < len(n); i++{
        if n[i] > max {
            max = n[i]
        }
    }
    return int(max - '0')
}
