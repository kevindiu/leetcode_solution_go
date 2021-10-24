// https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
    l := len(strs)
    if l == 0 {
        return ""
    }
    if l == 1 {
        return strs[0]
    }
    
    // loop through str 0
    for i, c := range strs[0] {
        matched := true
        ch := byte(c)
        
        // loop through remaining strs
        for j := 1; j < l; j++ {
            if i >= len(strs[j]) {
                return strs[0][:i]
            }
            
            if strs[j][i] != ch {
                matched = false
            }
        }
        
        if !matched {
            return strs[0][:i]
        }
    }
    
    return strs[0]
}
