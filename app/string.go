package app

func StringInSlice(itemToSearch string, list []string) bool {
	for _, item := range list {
		if item == itemToSearch {
			return true
		}
	}
	return false
}
