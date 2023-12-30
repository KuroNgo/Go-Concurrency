package main

import (
	_os "Learning-RxGo/os"
	_switch "Learning-RxGo/switch"
	"fmt"
)

func main() {
	DisplayMenu()
}

func DisplayMenu() {
	var num int

	fmt.Println("Vui lòng chọn kiến thức trong golang về Concurrency!")
	fmt.Println("1. Về Concurrency")
	fmt.Println("2. Về RxGO")

	fmt.Print("Input number: ")
	fmt.Scan(&num)

	switch num {
	case 1:
		_os.OperatingSystem()
		fmt.Println("1. Về Concurrency: ")
		_switch.Concurrency()

	case 2:
		_os.OperatingSystem()
		fmt.Println("2. Về RxGO")

	default:
		fmt.Println("Lựa chọn không hợp lệ. Hãy thử lại.")
		DisplayMenu()
	}

}
