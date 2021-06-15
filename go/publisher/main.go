package main

import (
	"fmt"
	"time"
	"sync"
	"strings"
)

type topicFilterFunc func(v interface{})bool

type subscriber struct {
    buffer   int
    ch chan interface{}
	tf topicFilterFunc
}

type publisher struct {
    timeout time.Duration
    subscribers []subscriber
}

func newPublisher(timeout time.Duration) *publisher {
    return &publisher{timeout: timeout}
}

func newSubsciber(buf int, tf topicFilterFunc) *subscriber {
	return &subscriber{buffer: buf, ch:make(chan interface{}, buf), tf: tf}
}

func (p *publisher)addSubscribe(sub subscriber) {
    p.subscribers = append(p.subscribers, sub)
}

func (p *publisher)publish(v interface{}) {
	var wg sync.WaitGroup
	for _, sub :=range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, v, &wg)
	}
	wg.Wait()
}
	

func (p *publisher) close() {
	for i, _ :=range p.subscribers {
		close(p.subscribers[i].ch)
		p.subscribers = append(p.subscribers[0:i], p.subscribers[i+1:]...)
	}
}

func (p *publisher)sendTopic(sub subscriber, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	if sub.tf == nil || !sub.tf(v) {
		fmt.Println("subsciber filted msg!")
        return
    }
	
	select {
		case sub.ch <- v:
			fmt.Println("publisher send done")
		case <-time.After(p.timeout):
			fmt.Println("time out")
	}
}

func main() {
	p := newPublisher(100 *time.Millisecond)
	defer p.close()

	s := newSubsciber(10, func(v interface{})bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "hello")
		}
		return false
	})

	p.addSubscribe(*s)

	p.publish("hello1")
	p.publish("hello2")
	p.publish("hello3")
	p.publish("golang3")

	go func ()  {
		for msg :=range s.ch {
			fmt.Printf("subsciber reciver: %s \n", msg)    
		}
	}()

	time.Sleep(time.Second *3)
}
