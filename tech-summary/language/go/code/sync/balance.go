// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// slides: https://talks.golang.org/2012/waza.slide#37
// code: https://talks.golang.org/2010/io/balance.go

// +build ignore,OMIT

package main

import (
	"container/heap"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const nRequester = 100
const nWorker = 10

var roundRobin = flag.Bool("r", false, "use round-robin scheduling")

// Simulation of some work: just sleep for a while and report how long.
func op() int {
	n := rand.Int63n(1e9)
	time.Sleep(time.Duration(nWorker * n))
	return int(n)
}

type Request struct {
	fn func() int
	c  chan int
}

// [Perry] requester simulate the call from external users
func requester(work chan Request) {
	c := make(chan int)
	for {
		time.Sleep(time.Duration(rand.Int63n(nWorker * 2e9)))
		// [Perry] send request
		work <- Request{op, c}
		// [Perry] wait for answer, result:= <-c, then further process on result
		<-c
	}
}

type Worker struct {
	i        int            // index in the loop
	requests chan Request   // work to do(a buffered channel)
	pending  int            // count of pending tasks
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests  // get request from load balancer
		req.c <- req.fn()    // call fn and send result to requester
		done <- w            // tell balancer we finished the job, send worker's pointer 
	}
}

// slice
type Pool []*Worker

// [Perry] why here using Pool not *Pool
// 1. Pool is not going to change
// 2.Passing by value often is cheaper in golang
// Go uses escape analysis to determine if variable can be safely
// allocated on function’s stack frame, which could be much cheaper
// then allocating variable on the heap. Passing by value simplifies
// escape analysis in Go and gives variable a better chance to be
// allocated on the stack.
// https://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/
func (p Pool) Len() int { return len(p) }

func (p Pool) Less(i, j int) bool {
	// [Perry]? safety check for i,j?  
	return p[i].pending < p[j].pending
}

func (p *Pool) Swap(i, j int) {
	a := *p
	a[i], a[j] = a[j], a[i]
	a[i].i = i
	a[j].i = j
}

func (p *Pool) Push(x interface{}) {
	a := *p
	n := len(a)
	a = a[0 : n+1]
	w := x.(*Worker)
	a[n] = w
	w.i = n
	*p = a
}

func (p *Pool) Pop() interface{} {
	a := *p
	*p = a[0 : len(a)-1]
	w := a[len(a)-1]
	w.i = -1 // for safety
	return w
}


type Balancer struct {
	pool Pool                // pool is the slice represent for workers, make(Pool, 0, nWorker)
	done chan *Worker        // make(chan *Worker, nWorker), when worker finished put in this channel, like a trigger
	i    int                 // i used for round robin
}

func NewBalancer() *Balancer {
	done := make(chan *Worker, nWorker)
	b := &Balancer{make(Pool, 0, nWorker), done, 0}
	for i := 0; i < nWorker; i++ {
		w := &Worker{requests: make(chan Request, nRequester)}
		heap.Push(&b.pool, w)
		go w.work(b.done)
	}
	return b
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work:
			b.dispatch(req)
		case w := <-b.done:
			b.completed(w)
		}
		b.print()
	}
}

func (b *Balancer) print() {
	sum := 0
	sumsq := 0
	for _, w := range b.pool {
		fmt.Printf("%d ", w.pending)
		sum += w.pending
		sumsq += w.pending * w.pending
	}
	avg := float64(sum) / float64(len(b.pool))
	variance := float64(sumsq)/float64(len(b.pool)) - avg*avg
	fmt.Printf(" %.2f %.2f\n", avg, variance)
}

// Send request to worker
func (b *Balancer) dispatch(req Request) {
	if *roundRobin {
		w := b.pool[b.i]
		w.requests <- req
		w.pending++
		b.i++
		if b.i >= len(b.pool) {
			b.i = 0
		}
		return
	}

	// get least loaded worker
	// [Perry] why  .(*Worker)
	// Tried with (*Worker)(), but got: cannot convert heap.Pop(&b.pool) (type interface {}) to type *Worker: need type assertion
	w := heap.Pop(&b.pool).(*Worker)
	// assign the task
	w.requests <- req
	// add one more in its queue
	w.pending++
	//	fmt.Printf("started %p; now %d\n", w, w.pending)
	// pub it back to heap
	heap.Push(&b.pool, w)
}

// job is complete; update heap
func (b *Balancer) completed(w *Worker) {
	if *roundRobin {
		w.pending--
		return
	}

	// one fewer in its queue
	w.pending--
	//	fmt.Printf("finished %p; now %d\n", w, w.pending)
	// remove it from heap
	heap.Remove(&b.pool, w.i)
	// adjust heap by push back
	heap.Push(&b.pool, w)
}

func main() {
	flag.Parse()
	// unbuffered channel, all request could go into loadbalancer one by one
	// due to code in requester => work <- Request{op, c}
	work := make(chan Request)
	for i := 0; i < nRequester; i++ {
		go requester(work)
	}
	NewBalancer().balance(work)
}
