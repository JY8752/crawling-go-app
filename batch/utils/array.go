package utils

func Contains(arr []string, item string) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func Remove(arr []string, item string) []string {
	var na []string
	for _, v := range arr {
		if v != item {
			na = append(na, v)
		}
	}
	return na
}
