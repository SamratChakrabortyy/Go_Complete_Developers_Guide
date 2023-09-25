package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxSleepMillis = 500

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	c := make(chan string)

	for i := 0; i < 5; i++ {
		go doSomething(i, r, c)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}

func doSomething(i int, r *rand.Rand, c chan string) {
	dur := time.Duration(r.Intn(500)) * time.Millisecond
	time.Sleep(dur)
	c <- fmt.Sprintln(i, "Slept for", dur)
}
