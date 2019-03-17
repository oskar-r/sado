package utility

func Contains(c string, s []string) bool {
	for _, v := range s {
		if c == v {
			return true
		}
	}
	return false
}
