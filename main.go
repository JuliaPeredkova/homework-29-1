package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		number := sc.Text()

		if number == "стоп" {
			return
		}

		resultNumber, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		fmt.Println("Ввод: ", resultNumber)
		fc := numberSquare(resultNumber)
		sc := multipicationByTwo(fc)
		fmt.Println("Произведение: ", <-sc)
	}

}

func numberSquare(n int) chan int {
	firstChan := make(chan int)

	go func() {
		firstChan <- n * n
	}()

	return firstChan
}

func multipicationByTwo(firstChan chan int) chan int {
	secondChan := make(chan int)
	nn := <-firstChan
	fmt.Println("Квадрат: ", nn)

	go func() {
		secondChan <- nn * 2
	}()

	return secondChan
}
