package main

import (
	"fmt"
	"time"
)

type stopWatch struct {
	st time.Time
}

func (s *stopWatch) start() {
	s.st = time.Now()
}

func (s *stopWatch) end() {
	end := time.Now()
	fmt.Printf("%f秒", end.Sub(s.st).Seconds())
}

func main() {
	sw := &stopWatch{}
	start(sw)
	end(sw)
}

func start(sw *stopWatch) {
	sw.start()
}

func end(sw *stopWatch) {
	sw.end()
}
