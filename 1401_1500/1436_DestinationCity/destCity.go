package destCity

func destCity(paths [][]string) string {
	pathLen := len(paths)
	placeMap := make(map[string]struct{}, pathLen)
	for i := 0; i < pathLen; i++ {
		placeMap[paths[i][0]] = struct{}{}
	}
	for i := 0; i < pathLen; i++ {
		if _, ok := placeMap[paths[i][1]]; !ok {
			return paths[i][1]
		}
	}
	return ""
}
