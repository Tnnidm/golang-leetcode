package guessNumber

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int

func guessNumber(n int) int {
	if n == 1 {
		return 1
	} else {
		begin, end := 1, n
		var mid int
		for {
			mid = (begin + end) / 2
			if guess(mid) == 0 {
				return mid
			} else if guess(mid) == -1 {
				end = mid - 1
			} else {
				begin = mid + 1
			}
		}
	}
}
