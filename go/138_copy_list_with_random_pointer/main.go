package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// code
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	mapNodeToNewNode := make(map[*Node]*Node)

	dummyHead := &Node{}

	iter := dummyHead
	for head != nil {
		if mapNodeToNewNode[head] == nil {
			mapNodeToNewNode[head] = &Node{
				Val: head.Val,
			}
		}

		if head.Next != nil {
			if mapNodeToNewNode[head.Next] == nil {
				mapNodeToNewNode[head.Next] = &Node{
					Val: head.Next.Val,
				}
			}
			mapNodeToNewNode[head].Next = mapNodeToNewNode[head.Next]
		}
		if head.Random != nil {
			if mapNodeToNewNode[head.Random] == nil {
				mapNodeToNewNode[head.Random] = &Node{
					Val: head.Random.Val,
				}
			}
			mapNodeToNewNode[head].Random = mapNodeToNewNode[head.Random]
		}

		iter.Next = mapNodeToNewNode[head]

		iter = iter.Next
		head = head.Next
	}

	return dummyHead.Next
}
