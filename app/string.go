package app

func StringPosInSlice(itemToSearch string, list []string) int {
	for i, item := range list {
		if item == itemToSearch {
			return i
		}
	}
	return -1
}
