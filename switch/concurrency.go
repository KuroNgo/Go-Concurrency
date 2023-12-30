package _switch

import (
	"Learning-RxGo/concurrency"
	_os "Learning-RxGo/os"
	"fmt"
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
		_os.OperatingSystem()
		fmt.Println("Về goroutine")
		concurrency.Goroutines()

	case 2:
		_os.OperatingSystem()
		fmt.Println("Về Channel")

	case 3:
		_os.OperatingSystem()
		fmt.Println("Về Select")

	case 4:
		_os.OperatingSystem()
		fmt.Println("Về Wait Group")

	case 5:
		_os.OperatingSystem()
		fmt.Println("Về tài liệu tham khảo")

	default:
		fmt.Println("Lựa chọn không hợp lệ. Vui lòng nhập lại")
		Concurrency()
	}
}
