// https://leetcode.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
    lolw := 0
    i := len(s)-1
    for ; i >= 0; i-- {
        if s[i] != ' ' {
            break
        }
    }
    for ; i>= 0; i-- {
        if s[i] == ' ' {
            break
        }
        lolw++
    }
    return lolw
}
