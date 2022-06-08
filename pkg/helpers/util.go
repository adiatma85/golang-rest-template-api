package helpers

import "strconv"

// Helper to change string to uint
func ConvertStringtoUint(id string) uint {
	typeUint64, _ := strconv.ParseUint(id, 10, 64)
	return uint(typeUint64)
}
