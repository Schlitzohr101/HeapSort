package main

import "fmt"

func main() {
	C := &Controller{}

	C.Menu()

	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("Error occured getting selection")
	}
	for i != 7 {
		method := C.choiceTaker(i)

		method()
		C.Menu()
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			fmt.Println("Error occured getting selection")
		}

	}
	fmt.Println("Thank you have a nice day")
}
