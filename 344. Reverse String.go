// https://leetcode.com/problems/reverse-string/
func reverseString(s []byte)  {
    for i := 0; i < len(s) / 2; i++ {
        lastIdx := len(s) - 1 -i
        s[i], s[lastIdx] = s[lastIdx], s[i]
    }
}
