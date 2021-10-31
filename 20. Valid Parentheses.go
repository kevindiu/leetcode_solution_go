// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
    l := len(s)
    if l == 0 {
        return true
    }
    if l % 2 != 0 {
        return false
    }
    
    stack := []byte{}
    for _, c1 := range s {
        c:= byte(c1)
        switch c {
            case '(', '{', '[':
                // push to the stack
                stack = append(stack, c)
            default:
                l := len(stack)
                if l == 0 {
                    return false
                }
                  
                // pop from the stack
                v := stack[l -1]
                if !match(v, c) {
                    return false
                }
                stack = stack[:l-1]
        }
    }    
    
    return len(stack) == 0
}

func match(c1, c2 byte) bool {
    switch c1 {
        case '(':
            return c2 == ')'
        case '{':
         return c2 == '}'
        case '[':
            return c2 == ']'
    }
    return false
}
