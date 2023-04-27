package helper

func UintOptional(val uint64, def uint64) uint64 {
	if val > 0 {
		return val
	}
	return def
}
