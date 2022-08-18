package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	balancer := NewBalancer()
	balancer.SetClient(&Client{name: "Client2"})
	balancer.SetClient(&Client{name: "Client1"})
	balancer.getClient().DO()
}

type Balancer struct {
	client []*Client
	size   int32
}

type Client struct {
	name string
}

func NewBalancer() *Balancer {
	balancer := &Balancer{client: []*Client{}}
	return balancer
}
func (bl *Balancer) SetClient(client *Client) *Balancer {
	bl.client = append(bl.client, client)
	bl.size += 1
	return bl
}

// 随机轮训
func (bl *Balancer) getClient() *Client {
	rand.Seed(time.Now().Unix())
	r := rand.Int31n(100)
	return bl.client[r%bl.size]
}

func (client *Client) DO() {
	fmt.Println("do client:", client.name)
}
