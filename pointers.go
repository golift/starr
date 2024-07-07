package starr

// Ptr returns a pointer to the provided "whatever".
func Ptr[P any](p P) *P {
	return &p
}

// True returns a pointer to a true boolean.
func True() *bool {
	return Ptr(true)
}

// False returns a pointer to a false boolean.
func False() *bool {
	return Ptr(false)
}

// String returns a pointer to a string.
// Deprecated: Use Ptr() function instead.
func String(s string) *string {
	return &s
}

// Int64 returns a pointer to the provided integer.
// Deprecated: Use Ptr() function instead.
func Int64(s int64) *int64 {
	return &s
}
