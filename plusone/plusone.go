package plusone

import "fmt"

//加一
func plusOne(digits []int) []int {
	l := len(digits)

	//如果长度为零则直接返回[1]
	if l == 0 {
		return []int{1}
	}

	//从切片尾部向前遍历
	for i := l-1; i>= 0; i-- {
		//判断是否为9 否则直接加1返回 是则将次位置设为0 指针向前移动
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	//如果切片中全是9 则在此返回 切片首位插入1
	return append([]int{1}, digits...)
}

func Test(){
	a := []int{9,9,9}
	fmt.Print(a)
	fmt.Print("+")
	fmt.Print(1)
	fmt.Print("=")
	b := plusOne(a)
	fmt.Println(b)
}