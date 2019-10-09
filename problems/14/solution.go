package leetcode150

type node struct {
	val  byte
	cnt  int
	next *node
}

func longestCommonPrefix(strs []string) string {
	head := &node{}
	for _, str := range strs {
		tmp := head
		for i := 0; i < len(str); i++ {
			if tmp.next == nil {
				tmp.next = &node{val: str[i], cnt: 1}
			} else {
				if tmp.next.val == str[i] {
					tmp.next.cnt++
				} else {
					break
				}
			}
			tmp = tmp.next
		}
	}

	var result []byte
	for head.next != nil {
		head = head.next
		if head.cnt == len(strs) {
			result = append(result, head.val)
		} else {
			break
		}
	}
	return string(result)
}
