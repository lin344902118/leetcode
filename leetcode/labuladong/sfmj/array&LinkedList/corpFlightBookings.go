package main

/*
1109. 航班预订统计
这里有 n 个航班，它们分别从 1 到 n 进行编号。
有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi]
意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。
*/

func corpFlightBookings(bookings [][]int, n int) []int {
	tmp := make([]int, n+1)
	for _, book := range bookings {
		tmp[book[0]-1] += book[2]
		tmp[book[1]] -= book[2]
	}
	for i := 1; i < len(tmp); i++ {
		tmp[i] += tmp[i-1]
	}
	return tmp[:n]
}

func main() {

}
