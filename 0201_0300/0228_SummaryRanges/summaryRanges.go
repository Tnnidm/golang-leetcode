package summaryRanges

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	numsLen := len(nums)
	if numsLen == 0 {
		return []string{}
	} else if numsLen == 1 {
		return []string{strconv.Itoa(nums[0])}
	} else {
		begin, end := nums[0], nums[0]
		result := []string{}
		for i := 0; i < numsLen-1; i++ {
			if nums[i+1] == nums[i]+1 {
				end = nums[i+1]
			} else {
				if begin == end {
					result = append(result, strconv.Itoa(end))
				} else {
					result = append(result, strconv.Itoa(begin)+"->"+strconv.Itoa(end))
				}
				begin = nums[i+1]
				end = nums[i+1]
			}
		}
		if begin == end {
			result = append(result, strconv.Itoa(end))
		} else {
			result = append(result, strconv.Itoa(begin)+"->"+strconv.Itoa(end))
		}
		return result
	}
}

// // 更简洁的写法
// func summaryRanges(nums []int) (ret []string) {
// 	for i := 0; i < len(nums); {
// 		left := i
// 		for i++; i < len(nums) && nums[i-1]+1 == nums[i]; i++ {
// 		}
// 		tmp := strconv.Itoa(nums[left])
// 		if left < i-1 {
// 			tmp += fmt.Sprintf("->%d", nums[i-1])
// 		}
// 		ret = append(ret, tmp)
// 	}
// 	return ret
// }
