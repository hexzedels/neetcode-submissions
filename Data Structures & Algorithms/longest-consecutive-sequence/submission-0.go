type node struct {
    cur int
    prev *node
    next *node
}

func longestConsecutive(nums []int) int {
    arr := make(map[int]*node, len(nums))

    for _, n := range nums {        
        curNode := &node{
            cur: n,
        }

        if prev, ok := arr[n-1]; ok {
            curNode.prev = prev
            prev.next = curNode
        }

        if next, ok := arr[n+1]; ok {
            curNode.next = next
            next.prev = curNode
        }

        arr[n] = curNode
    }

    var ind []int
    var maxL int

for i := 0; len(arr) > 0; i++{
    node, ok := arr[nums[i]]
    if ok {
        ind = append(ind, node.cur)
        prevNode := node.prev
        nextNode := node.next
        for prevNode != nil {
            ind = append(ind, prevNode.cur)
            prevNode = prevNode.prev
        }

        for nextNode != nil {
            ind = append(ind, nextNode.cur)
            nextNode = nextNode.next
        }

        maxL = max(maxL, len(ind))
        for _, in := range ind {
            delete(arr, in)
        }
        ind = []int{}
    }
}

return maxL
}