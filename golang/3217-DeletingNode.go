package main
type ListNode struct{
  Val int
  Next  *ListNode 
}
func ModifiedList(nums []int, head *ListNode) *ListNode {
  numsMap := make(map[int]bool)

  for _, v := range nums{
    numsMap[v] = true 
  }

  iNode :=   &ListNode{}
  iNode.Next = head 
  pNode := iNode
  cNode := head

  for cNode != nil{
    if numsMap[cNode.Val]{
      pNode.Next = cNode.Next
    }else{
      pNode = cNode 
    }
    cNode = cNode.Next
  }
  return iNode.next
}
