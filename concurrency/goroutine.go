package concurrency

import (
	"fmt"
	"os"
	"os/exec"
)

func Goroutines() {
	var num int

	fmt.Println("Các kiến thức liên quan đến goroutine")
	fmt.Println("======================================================================")
	fmt.Println("1. Khái niệm")
	fmt.Println("2. So sánh về OS Thread")
	fmt.Println("3. Cách khai báo một Goroutine")
	fmt.Println("4. Sử dụng Goroutine với anonymous function Golang")
	fmt.Println("======================================================================")
	fmt.Print("Vui lòng chọn số kiến thức mà bạn quan tâm: ")
	fmt.Scan(&num)

	switch num {
	case 1:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("Kiến thức về Goroutine")
		fmt.Println("======================================================")
		fmt.Println("1. Khái niệm")

		// hàm ở đây
	case 2:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("Kiến thức về Goroutine")
		fmt.Println("======================================================")
		fmt.Println("2. So sánh về OS Thread")

		// hàm ở đây

	case 3:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("Kiến thức về Goroutine")
		fmt.Println("======================================================")
		fmt.Println("3. Cách khai báo một Goroutine")

		// hàm ở đây

	case 4:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("Kiến thức về Goroutine")
		fmt.Println("======================================================")
		fmt.Println("4.Sử dụng Goroutine với anonymous function Golang")

		// hàm ở đây
	}
}
