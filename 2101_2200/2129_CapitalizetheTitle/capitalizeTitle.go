package capitalizeTitle

func capitalizeTitle(title string) string {
	titleLen := len(title)
	begin, end := -1, -1
	titleByte := []byte(title)
	capitalize := byte('a' - 'A')
	for end != titleLen {
		begin = end + 1
		end++
		for end < titleLen && titleByte[end] != ' ' {
			end++
		}
		if end-begin <= 2 { // 长度小于等于2，全部转小写
			for i := begin; i < end; i++ {
				if titleByte[i] >= 'A' && titleByte[i] <= 'Z' {
					titleByte[i] += capitalize
				}
			}
		} else { // 长度大于2，除了第一个字母大写，其余全部转小写
			if titleByte[begin] >= 'a' && titleByte[begin] <= 'z' {
				titleByte[begin] -= capitalize
			}
			for i := begin + 1; i < end; i++ {
				if titleByte[i] >= 'A' && titleByte[i] <= 'Z' {
					titleByte[i] += capitalize
				}
			}
		}
	}
	return string(titleByte)
}
