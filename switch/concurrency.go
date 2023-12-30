package _switch

import (
	"Learning-RxGo/concurrency"
	"fmt"
	"os"
	"os/exec"
)

func Concurrency() {
	var num int

	fmt.Println("Các kiến thức về Concurrency")
	fmt.Println("======================================================================")
	fmt.Println("1. Goroutine")
	fmt.Println("2. Channel")
	fmt.Println("3. Select")
	fmt.Println("4. Wait Group")
	fmt.Println("5. Tài liệu tham khảo")
	fmt.Println("======================================================================")
	fmt.Print("Vui lòng chọn kiến thức mà bạn quan tâm: ")
	fmt.Scan(&num)

	switch num {
	case 1:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("Về goroutine")
		concurrency.Goroutines()

	case 2:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("Về Channel")

	case 3:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("Về Select")

	case 4:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("Về Wait Group")

	case 5:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("Về tài liệu tham khảo")
	}
}
