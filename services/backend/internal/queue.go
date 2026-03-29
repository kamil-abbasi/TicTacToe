package internal

import "log"

type Queue struct {
	enqueue chan *Client
}

func NewQueue() *Queue {
	return &Queue{
		enqueue: make(chan *Client),
	}
}

func (q *Queue) Enqueue(client *Client) {
	q.enqueue <- client
	log.Printf("client %v enqueued", client.Name())
}

func (q *Queue) Run() {
	for {
		client1 := <-q.enqueue
		log.Printf("client %v dequeued", client1.Name())
		client2 := <-q.enqueue
		log.Printf("client %v dequeued", client2.Name())

		room := NewRoom(client1, client2)
		go room.Run()
		log.Printf("Room with clients: %v, %v started", client1.Name(), client2.Name())
	}
}
