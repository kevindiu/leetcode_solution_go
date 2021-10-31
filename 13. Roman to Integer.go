// https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) int {
    svs := make([]int, int('Z'-'A'))
    svs['I'-'A'] = 1
    svs['V'-'A'] = 5
    svs['X'-'A'] = 10
    svs['L'-'A'] = 50
    svs['C'-'A'] = 100
    svs['D'-'A'] = 500
    svs['M'-'A'] = 1000
    
    lastValue := svs[int(s[len(s) - 1]-'A')]
    cnt := lastValue
    for i:=len(s) - 2; i>=0; i-- {
        v := svs[s[i]-'A']
        if lastValue > v {
            cnt -= v
        } else {
            cnt += v
            lastValue = v
        }
    }
    return cnt
}
