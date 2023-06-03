package main

/*
  有效的数独
请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）


注意：

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
空白格用 '.' 表示。
*/

func isValidSudoku(board [][]byte) bool {
	// 判断横竖，九宫格都没有重复
	for i := 0; i < len(board); i++ {
		arr1 := make([]byte, 0, 9)
		arr2 := make([]byte, 0, 9)
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				arr1 = append(arr1, board[i][j])
			}
			if board[j][i] != '.' {
				arr2 = append(arr2, board[j][i])
			}
		}
		if isRepeat(arr1) {
			return false
		}
		if isRepeat(arr2) {
			return false
		}
	}
	for i := 1; i < len(board); i += 3 {
		for j := 1; j < len(board[i]); j += 3 {
			arr := make([]byte, 0, 9)
			for k := i - 1; k < i+2; k++ {
				for l := j - 1; l < j+2; l++ {
					if board[k][l] != '.' {
						arr = append(arr, board[k][l])
					}
				}
			}
			if isRepeat(arr) {
				return false
			}
		}
	}
	return true
}

func isRepeat(arr []byte) bool {
	cache := make(map[byte]struct{})
	for _, num := range arr {
		_, exist := cache[num]
		if exist {
			return true
		}
		cache[num] = struct{}{}
	}
	return false
}

func main() {

}
