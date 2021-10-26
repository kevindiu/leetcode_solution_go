// https://leetcode.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
    lolw := 0
    for i := len(s)-1; i >= 0; i-- {
        if s[i] == ' ' {
            if lolw == 0 {
                continue
            }
            return lolw
        }
        lolw++
    }
    return lolw
}
