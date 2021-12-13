package gcdOfStrings

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	for {
		difference := len(str1) - len(str2)
		if difference == 0 {
			return str1
		} else if difference > 0 {
			str1 = str1[len(str2):]
		} else {
			str2 = str2[len(str1):]
		}
	}
}
