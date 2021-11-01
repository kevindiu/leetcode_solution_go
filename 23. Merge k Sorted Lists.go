// https://leetcode.com/problems/merge-k-sorted-lists/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    l := len(lists)
    if l == 0 {
        return nil
    } 
    if l == 1 {
        return lists[0]
    }
    
    h := new(IntHeap)
    for i :=0; i<l; i++ {
        if lists[i] != nil {
            h.Push(node {
                Val: lists[i].Val,
                Idx: i,
            })
            lists[i] = lists[i].Next
        }
    }
    
    if h.Len() == 0 {
        return nil
    }
    heap.Init(h)
      
    n := heap.Pop(h).(node)
    result := &ListNode{
        Val: n.Val,
    }
    
    if lists[n.Idx] != nil {
        heap.Push(h, node {
            Val: lists[n.Idx].Val,
            Idx: n.Idx,
        })
        lists[n.Idx] = lists[n.Idx].Next
    }
    
    cur := result
    for h.Len() > 0 {
        n := heap.Pop(h).(node)
        cur.Next = &ListNode{Val: n.Val}
        cur = cur.Next
        
        if lists[n.Idx] != nil {
            heap.Push(h, node {
               Val: lists[n.Idx].Val,
               Idx: n.Idx,
            })
            lists[n.Idx] = lists[n.Idx].Next
         }
    }
    
    return result
}

type node struct {
    Val int
    Idx int
}

type IntHeap []node

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(node))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
