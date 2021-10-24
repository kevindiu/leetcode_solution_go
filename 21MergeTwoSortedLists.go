// https://leetcode.com/problems/merge-two-sorted-lists/submissions/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }
    
    var (
        res *ListNode
        ln1 = l1
        ln2 = l2
    )
    popLn1 := func() int {
        v := ln1.Val
        ln1 = ln1.Next
        return v
    }
    popLn2 := func() int {
        v := ln2.Val
        ln2 = ln2.Next
        return v
    }

    if l1.Val > l2.Val {
        res = &ListNode { popLn2(), nil }
    } else {
        res = &ListNode { popLn1(), nil }
    }
    
    cur := res
    append := func(val int) {
        cur.Next = &ListNode {val, nil}
        cur = cur.Next
    }
    
    for {
        if ln1 == nil && ln2 == nil {
            return res
        }
        
        if ln1 == nil {
            // append remaining v2 to result
            for ln2 != nil {
                append(popLn2())
            }
            return res
        }

        if ln2 == nil {
            // append remaining v1 to result
            for ln1 != nil {
                append(popLn1())
            }
            return res
        }
        
        if ln1.Val > ln2.Val {
            append(popLn2())
        } else if ln2.Val > ln1.Val {
            append(popLn1())
        } else {
            append(popLn1())
            append(popLn2())
        }
    }
}

