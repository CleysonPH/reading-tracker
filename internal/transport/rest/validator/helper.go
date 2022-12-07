package validator

func stringIn(value string, opt ...string) bool {
	for _, item := range opt {
		if item == value {
			return true
		}
	}
	return false
}
