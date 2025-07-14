/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	value := head
	stringDecimal := ""

	for value != nil {
		temp := strconv.Itoa(value.Val)
		stringDecimal = temp + stringDecimal

		if value.Next == nil {
			value = nil
		} else {
			next := value.Next
			value = next
		}
	}

	res := 0
	for i := 0; i < len(stringDecimal); i++ {
		b := stringDecimal[i]
		if b == '1' {
			res += int(math.Pow(2, float64(i)))
		}
	}

	return res
}