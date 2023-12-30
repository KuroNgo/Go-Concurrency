package concurrency

import (
	_os "Learning-RxGo/os"
	"fmt"
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
		_os.OperatingSystem()
		fmt.Println("Kiến thức về Goroutine")
		fmt.Println("======================================================")
		fmt.Println("1. Khái niệm")

		// hàm ở đây
		TheNationOfGoroutine()
	case 2:
		_os.OperatingSystem()
		fmt.Println("Kiến thức về Goroutine")
		fmt.Println("======================================================")
		fmt.Println("2. So sánh về OS Thread")

		// hàm ở đây
		CompareToOSThread()
	case 3:
		_os.OperatingSystem()
		fmt.Println("Kiến thức về Goroutine")
		fmt.Println("======================================================")
		fmt.Println("3. Cách khai báo một Goroutine")

		// hàm ở đây

	case 4:
		_os.OperatingSystem()
		fmt.Println("Kiến thức về Goroutine")
		fmt.Println("======================================================")
		fmt.Println("4.Sử dụng Goroutine với anonymous function Golang")

		// hàm ở đây

	default:
		fmt.Println("Lựa chọn không hợp lệ. Vui lòng nhập lại")
		Goroutines()
	}
}

func TheNationOfGoroutine() {
	fmt.Println("Trong bất kỳ một chương trình Golang đều tồn tại ít nhất một Goroutine, gọi là main Goroutine")
	fmt.Println("Nếu main goroutines này kết thúc, toàn bộ các goroutine khác trong chương trình đều bị dừng và kết thúc ngay.")
	fmt.Println("Goroutine bản chất là một lightweight execution thread (luồng thực thi gọn nhẹ)")
	fmt.Println("Vì thế việc sử dụng các Goroutines trong Golang có chi phí cực kỳ thấp so với cách sử dụng các Thread truyền thống (OS thread)")
}

func CompareToOSThread() {
	fmt.Println("Goroutines có kích thước nhỏ hơn rất đáng kể so với Thread. " +
		"\nGoroutines sử dụng 2KB memory stack, trong khi đó các OS Thread lên đến 2MB." +
		"\nGoroutines có thể linh động tăng giảm bộ nhớ sử dụng trong khi đó OS Thread là cố định." +
		"\nMột chương trình Golang có thể có hàng trăm nghìn Goroutine trong khi Thread chỉ được vài trăm đến hàng nghìn." +
		"\nGoroutines có thời gian khởi động nhanh hơn Thread." +
		"\nGoroutines có thể giao tiếp an toàn với nhau thông qua các kênh trong Golang (Channel)." +
		"\nCác channel hỗ trợ mutex lock vì thế tránh được các lỗi liên quan tới cùng ghi và đọc lên vùng dữ liệu chia sẻ (data race)." +
		"\nGoroutines có thể được ánh xạ và hoạt động trên nhiều OS threads thay vì là ánh xạ 1:1.")
}
