package isIsomorphic

func isIsomorphic(s string, t string) bool {
	stMap := make(map[byte]byte)
	tsMap := make(map[byte]byte)
	for i := range s {
		st, ok_st := stMap[s[i]]
		ts, ok_ts := tsMap[t[i]]
		if ok_st {
			if ok_ts {
				if tsMap[st] != ts {
					return false
				}
			} else {
				return false
			}
		} else {
			if ok_ts {
				return false
			} else {
				stMap[s[i]] = t[i]
				tsMap[t[i]] = s[i]
			}
		}
	}
	return true
}
