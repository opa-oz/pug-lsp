package utils

// PtrBool returns a pointer to a bool value.
//
// It takes a bool input and returns a pointer to it. This allows passing and modifying
// bool values indirectly through the returned pointer.
//
// Parameters:
//
//	input bool - The boolean value to be referenced by the returned pointer.
//
// Returns:
//
//	*bool - A pointer to the input boolean value.
func PtrBool(input bool) *bool {
	return &input
}

func PtrIsTrue(v *bool) bool {
	return v != nil && *v == true
}
