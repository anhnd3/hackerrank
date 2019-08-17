package leftrotate

// RotLeft ...
func RotLeft(a []int32, d int32) []int32 {
	return append(a[d:], a[:d]...)
}
