package exercises

type Queue struct {
	data               []interface{}
	head, tail, length int
}

func NewQueue(size int) *Queue {
	data := make([]interface{}, size)
	return &Queue{
		data:   data,
		length: 0,
	}
}

func (q *Queue) EnQueue(data interface{}) bool {
	if q.length == len(q.data) {
		return false
	}

	if q.length == 0 {
		q.head = 0
		q.tail = 0
	} else {
		if t := q.tail + 1; t == len(q.data) {
			q.tail = 0
		} else {
			q.tail = t
		}
	}
	q.data[q.tail] = data
	q.length++
	return true
}

func (q *Queue) DeQueue() (data interface{}, ok bool) {
	if q.length == 0 {
		return nil, false
	}

	data = q.data[q.head]
	if t := q.head + 1; t == len(q.data) {
		q.head = 0
	} else {
		q.head = t
	}
	q.length--
	return data, true
}

func (q *Queue) Length() int {
	return q.length
}
