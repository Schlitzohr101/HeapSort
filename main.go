//William Murray
//John Miner
package main

import "fmt"

func main() {
	//Initialize Controller, set as pointer
	C := &Controller{}
	C.Menu()

	//Take in int for menu choice
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("Error occured getting selection")
	}
	for i != 7 {
		//Use Closure to get method chosen to run
		method := C.choiceTaker(i)

		//Run method
		method()
		//Repeat until finished
		C.Menu()
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			fmt.Println("Error occured getting selection")
		}

	}
	fmt.Println("Thank you have a nice day")
}
