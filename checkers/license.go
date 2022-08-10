package checkers

func Licence(key string) bool {
	if len(key) != 24 {
		return false
	}
	x := 1337
	for i, c := range key {
		if i == 4 || i == 9 || i == 14 || i == 19 {
			if c != '-' {
				return false
			}
		}
		x += int(c)
	}
	if x == 2872 {
		return true
	}
	return false
}
