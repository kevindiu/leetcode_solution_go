// https://leetcode.com/problems/container-with-most-water/

func maxArea(height []int) int {
    max := 0
    i, j:= 0, len(height)-1
    for i < j {
        max = maxInt(max, (j - i) * minInt(height[i], height[j]))
        if height[i] > height[j] {
            j--
        } else {
            i ++
        }
    }
    return max
}

func minInt(n1, n2 int) int {
    if n1 < n2 {
        return n1
    }
    return n2
}

func maxInt(n1, n2 int) int {
    if n1 > n2 {
        return n1
    }
    return n2
}
