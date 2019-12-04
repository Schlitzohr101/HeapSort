package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Controller struct {
	heap Tree
}

func (C Controller) Menu() {
	fmt.Print("MENU \n\n")
	fmt.Println("1) Create Empty heap")
	fmt.Println("2) Enter Integers to be used as heap")
	fmt.Println("3) Add to current heap")
	fmt.Println("4) Pop element from heap")
	fmt.Println("5) Sort the heap")
	fmt.Println("6) Create random array and examine")
	fmt.Println("7) Quit")
}

func (C Controller) choiceTaker(x int) func() {
	m := map[int]func(){
		1: C.option1,
		2: C.option2,
	}
	return m[x]
}

//generates a random int between the min and max provided
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (C Controller) option1() {
	fmt.Println("Creating a Empty Heap")
	ar := make([]int64, 0)
	C.heap.buildHeap(ar)
	C.heap.print()
}

func (C Controller) option2() {
	fmt.Println("Enter the Integers to be in the Heap followed by ',' i.e. : 1,2,3,...\n:")
	input := bufio.NewReader(os.Stdin)
	reply, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("reading in the Integers has failed unexpectedly")
	}
	reply = reply[:len(reply)-1]

	strs := strings.Split(reply, ",")
	ar := make([]int64, len(strs))
	for i := 0; i < len(strs); i++ {
		num, _ := strconv.Atoi(strs[i])
		ar[i] = int64(num)
	}
	C.heap.buildHeap(ar)
	C.heap.print()
}
