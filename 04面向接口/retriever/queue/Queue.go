package queue

/*
	interface{}  可以接受任何类型的对象值
*/
type QueuePlus []interface{}

// 向队列中加入一个元素
// 		e.g.  p.Push(123)
func (q *QueuePlus) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *QueuePlus) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *QueuePlus) IsEmpty() bool {
	return len(*q) == 0
}
