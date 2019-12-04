package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Controller struct {
	ar []int
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
	fmt.Print(": ")
}

func (C *Controller) choiceTaker(x int) func() {
	m := map[int]func(){
		1: C.option1,
		2: C.option2,
		3: C.option3,
		4: C.option4,
		5: C.option5,
		6: C.option6,
	}
	return m[x]
}

//generates a random int between the min and max provided
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func (C *Controller) option1() {
	fmt.Println("Creating a Empty Heap")
	C.ar = make([]int, 0)
	buildHeap(C.ar)
	fmt.Println(C.ar)
	// return C.ar
}

func (C *Controller) option2() {
	fmt.Print("Enter the Integers to be in the Heap followed by ',' i.e. : 1,2,3,...\n: ")
	input := bufio.NewReader(os.Stdin)
	reply, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("reading in the Integers has failed unexpectedly")
	}
	reply = reply[:len(reply)-1]

	strs := strings.Split(reply, ",")
	C.ar = make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		num, _ := strconv.Atoi(strs[i])
		C.ar[i] = num
	}
	buildHeap(C.ar)
	fmt.Println(C.ar)
}

func (C *Controller) option3() {

	fmt.Print("Enter a number to insert into the heap\n: ")
	input := bufio.NewReader(os.Stdin)
	reply, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("Failed reading in the number")
	}
	reply = reply[:len(reply)-1]

	num, _ := strconv.Atoi(reply)

	C.ar = append(C.ar, num)

	buildHeap(C.ar)
	fmt.Println("The heap after the insertion:")
	fmt.Println(C.ar)
	// return C.ar
}

func (C *Controller) option4() {
	if len(C.ar) == 0 {
		C.ar = make([]int, 0)
	} else {
		C.ar[0] = C.ar[len(C.ar)-1]
		C.ar = C.ar[:len(C.ar)-1]
	}
	buildHeap(C.ar)
	fmt.Println("Heap after pop")
	fmt.Println(C.ar)
}

func (C *Controller) option5() {
	ar := make([]int, len(C.ar))
	for i := 0; i < len(C.ar); i++ {
		ar[i] = C.ar[i]
	}
	fmt.Println("Heap before sort")
	fmt.Println(ar)
	fmt.Println("Heap after sort")
	heapSort(ar)
	fmt.Println(ar)
}

func (C *Controller) option6() {
	fmt.Print("Enter a size for the random array\n:")
	var size int
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		fmt.Println("Error occured getting size")
	}
	fmt.Println("Creating Array of size ", size)
	ar := make([]int, size)
	ar1 := make([]int, size)
	for i := 0; i < size; i++ {
		num := randInt(-1000, 1000)
		ar[i] = num
		ar1[i] = num
	}
	fmt.Println("Sorting with HeapSort")
	t1 := time.Now()
	heapSort(ar)
	t2 := time.Now()
	heapTime := t2.Sub(t1)
	fmt.Println("Elapsed time: ", heapTime)

	fmt.Println("Sorting with built in sort")
	t1 = time.Now()
	sort.Ints(ar1)
	t2 = time.Now()
	sortTime := t2.Sub(t1)
	fmt.Println("Elapsed time: ", sortTime)

	if heapTime-sortTime < 0 {
		fmt.Println("Our heapsort was faster by: ", sortTime-heapTime)
	} else {
		fmt.Println("The built in sort was faster by: ", heapTime-sortTime)
	}

}

func heapifyMin(ar []int, n int, i int) {
	smallest := i // Initialize largest as root
	l := 2*i + 1  // left = 2i + 1
	r := 2*i + 2  // right = 2i + 2

	// If left child is larger than root
	if l < n && ar[l] < ar[smallest] {
		smallest = l
	}

	// If right child is larger than largest so far
	if r < n && ar[r] < ar[smallest] {
		smallest = r
	}

	// If largest is not root
	if smallest != i {
		swap := ar[i]
		ar[i] = ar[smallest]
		ar[smallest] = swap

		// Recursively heapify the affected sub-tree
		heapifyMin(ar, n, smallest)
	}

}

func buildHeap(ar []int) {
	for index := len(ar) / 2; index >= 0; index-- {
		heapifyMin(ar, len(ar), index)
	}
}

func heapSort(ar []int) {
	buildHeap(ar)
	for i := len(ar) - 1; i >= 1; i-- {
		swap := ar[0]
		ar[0] = ar[i]
		ar[i] = swap
		ar = ar[:len(ar)-1]
		heapifyMin(ar, len(ar), 0)
	}
}
