package main

import (
	f "fmt"
	p "github.com/vuhoangphuc11/vhp-golang-campaign-4/pkg"
	"math/rand"
	"sync"
	t "time"
)

var mu sync.Mutex
var count = 0

func process(n int) {
	for i := 0; i < 10; i++ {
		t.Sleep(t.Duration(rand.Int31n(2)) * t.Millisecond)
		mu.Lock() // lock
		temp := count
		temp++
		t.Sleep(t.Duration(rand.Int31n(2)) * t.Millisecond)
		count = temp
		mu.Unlock() // unlock
	}
	f.Printf("Count after i= %v Count: %v\n", n, count)
}

func main() {

	//TODO Generic
	arr1 := []string{"foo", "bar", "baz"}
	f.Println(p.Index(arr1, "golang"))

	arr2 := []string{"a", "b", "c"}
	resultArr2 := p.Map(arr2, func(v string) string { return v + v })
	f.Println(resultArr2)

	//VHP Go goroutines
	go p.Numbers()
	go p.Alphabets()
	t.Sleep(3000 * t.Millisecond)
	f.Println("==/VU-HOANG-PHUC/==")

	//VHP Go channel with routine
	pipe1 := make(chan string)
	go func() {
		pipe1 <- "Tai Xiu"
	}()
	receiver1 := pipe1
	f.Println(receiver1)

	//VHP Go buffer channel
	pipe2 := make(chan string, 2)
	go func() {
		pipe2 <- "water"
		pipe2 <- "water"
		close(pipe2)
	}()
	for receiver := range pipe2 {
		f.Println(receiver)
	}

	//VHP Mutex
	for i := 1; i < 4; i++ {
		go process(i)
	}
	t.Sleep(10 * t.Second)
	f.Println("Final Count:", count)
}
