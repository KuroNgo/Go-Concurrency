package main

import (
	_switch "Learning-RxGo/switch"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var num int

	fmt.Println("Vui lòng chọn kiến thức trong golang về Concurrency!")
	fmt.Println("1. Về Concurrency")
	fmt.Println("2. Về RxGO")

	fmt.Print("Input number: ")
	fmt.Scan(&num)

	switch num {
	case 1:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("1. Về Concurrency: ")
		_switch.Concurrency()

	case 2:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("2. Về RxGO")
	}
}
