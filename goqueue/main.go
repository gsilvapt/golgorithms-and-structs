package goqueue

import errors

type Queue struct {
	front    int
	rear     int
	Q        []int
	capacity int
}

func NewQueue(cap int) *Queue {
	return &Queue{front: 0, rear: 0, Q: make([]int, cap), capacity: cap}
}

func (q *Queue) Enqueue(i int) (Error, *Queue) {
	if q.capacity == len(q.Q) {
		errors.New("Queue full.")
	}

	q.Q = append(q.Q, i)
	q.rear = q.rear + 1%q.capacity
	return nil, q
}

func (q *Queue) Dequeue() *Queue {
	q.Q = q.Q[q.front:]
	q.rear = q.rear % q.capacity
	return q
}
