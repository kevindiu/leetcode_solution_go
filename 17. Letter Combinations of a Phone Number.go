// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    
    letterMap := map[string][]string {
        "2": []string{"a","b","c"},
        "3": []string{"d","e","f"},
        "4": []string{"g","h","i"},
        "5": []string{"j","k","l"},
        "6": []string{"m","n","o"},
        "7": []string{"p","q","r","s"},
        "8": []string{"t","u","v"},
        "9": []string{"w","x","y","z"},
    }

    ans := letterMap[string(digits[0])]
    
    for i := 1; i < len(digits); i++ {
        tmp := make([]string, 0)
        d := string(digits[i])
        letters := letterMap[d]
        
        for j := 0; j < len(ans); j++ {            
            for _, l := range letters {
                tmp = append(tmp, ans[j]+l)
            }
        }
        ans = tmp
    }
    
    return ans
}

