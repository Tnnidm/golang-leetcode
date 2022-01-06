package orderedStream

type OrderedStream struct {
	ivMap map[int]string
	ptr   int
}

func Constructor(n int) OrderedStream {
	os := OrderedStream{}
	os.ivMap = make(map[int]string, n)
	os.ptr = 1
	return os
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.ivMap[idKey] = value
	result := []string{}
	if this.ptr == idKey {
		ok := true
		for ok {
			result = append(result, value)
			this.ptr++
			value, ok = this.ivMap[this.ptr]
		}
	}
	return result
}
