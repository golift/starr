package starr

// String returns a pointer to a string.
func String(s string) *string {
	return &s
}

// True returns a pointer to a true boolean.
func True() *bool {
	s := true
	return &s
}

// False returns a pointer to a false boolean.
func False() *bool {
	s := false
	return &s
}

// Int64 returns a pointer to the provided integer.
func Int64(s int64) *int64 {
	return &s
}
