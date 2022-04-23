package MyCalendar

// 二分搜索做法
// 思路是维护一个以start排序的数组
// 有新的book时，通过二分查找找到start插入的位置
// 然后检查和前一个和后一个有没有重复
type MyCalendar struct {
	books [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{[][2]int{{1000000000, 1000000001}}}
}

func (myCalendar *MyCalendar) Book(start int, end int) bool {
	bookLen := len(myCalendar.books)
	// 二分查找新的book插入的位置
	left, right := 0, bookLen-1
	for left < right {
		mid := left + (right-left)>>1
		if myCalendar.books[mid][0] == start {
			return false
		} else if myCalendar.books[mid][0] > start {
			right = mid
		} else {
			left = mid + 1
		}
	}
	// 检查前一个book的end是不是在新的这个之前
	if right-1 >= 0 && myCalendar.books[right-1][1] > start {
		return false
	}
	// 检查后一个book的start在新的一个end之后
	if right < bookLen && end > myCalendar.books[right][0] {
		return false
	}
	// 都通过，就添加这个book，返回true
	// 这里插入新元素，不用切片操作，而使用在尾部插入后冒泡上去，可以节省很多内存！
	myCalendar.books = append(myCalendar.books, [2]int{start, end})
	for i := bookLen; i > right; i-- {
		myCalendar.books[i], myCalendar.books[i-1] = myCalendar.books[i-1], myCalendar.books[i]
	}
	return true
}
