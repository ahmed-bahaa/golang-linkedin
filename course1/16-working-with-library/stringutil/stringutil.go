package stringutil

func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

func FullNameNaked(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}
