package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Ring struct {
	current          int
	firstPartAnswer  int
	secondPartAnswer int
}

func newRing() *Ring {
	r := new(Ring)
	r.current = 50
	return r
}

func (r *Ring) moveL(count int) {
	for _ = range count {
		r.sub()
	}
	if r.current == 0 {
		r.firstPartAnswer++
	}
}
func (r *Ring) moveR(count int) {
	for _ = range count {
		r.add()
	}
	if r.current == 0 {
		r.firstPartAnswer++
	}
}
func (r *Ring) add() {
	if r.current+1 > 99 {
		r.current = 0
		r.secondPartAnswer = r.secondPartAnswer + 1
		return
	} else {
		r.current++
	}
}
func (r *Ring) sub() {
	if r.current-1 < 0 {
		r.current = 99
		return
	}
	r.current--

	if r.current == 0 {
		r.secondPartAnswer = r.secondPartAnswer + 1
	}
}

func (r *Ring) isZero() bool {
	return r.current == 0
}

func main() {
	r := newRing()
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	str := bufio.NewScanner(f)
	for str.Scan() {

		val := str.Text()
		direction := val[0]
		count, err := strconv.Atoi(val[1:])
		if err != nil {
			panic(err)
		}
		
		if direction == 'L' {
			r.moveL(count)
			continue
		} else if direction == 'R' {
			r.moveR(count)
		}

	}

	fmt.Println(r.firstPartAnswer)
	fmt.Println(r.secondPartAnswer)
}
