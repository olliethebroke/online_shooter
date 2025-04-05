package utils

import "strconv"

// StringToInt32 converts string with number
// to the integer typed number.
//
// Accepts string to convert as an argument.
//
// Returns converted integer and error if
// it exists, otherwise nil.
func StringToInt32(str string) (int32, error) {
	num, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(num), nil
}

// StringToInt64 converts string with number
// to the integer typed number.
//
// Accepts string to convert as an argument.
//
// Returns converted integer and error if
// it exists, otherwise nil.
func StringToInt64(str string) (int64, error) {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

// StringToFloat converts string with number
// to the float32 typed number.
//
// Accepts string to convert as an argument.
//
// Returns converted float64 and error if
// it exists, otherwise nil.
func StringToFloat(str string) (float32, error) {
	num, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}
	return float32(num), nil
}
