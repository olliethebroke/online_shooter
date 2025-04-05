package utils

import (
	"errors"
	"fmt"
	"os"
)

// GetIntEnvVar gets an integer variable from the
// process environment.
//
// Accepts a name of the variable.
//
// Returns the variable's value and an error if it exists,
// otherwise nil.
func GetIntEnvVar(name string) (int32, error) {
	// get the var from the environment
	envStr := os.Getenv(name)
	if len(envStr) == 0 {
		return 0, errors.New(fmt.Sprintf("failed to get env variable: %s", name))
	}

	// parse string to int
	envInt, err := StringToInt32(envStr)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("failed to parse env variable: %s", name))
	}
	return envInt, nil
}

// GetFloatEnvVar gets a float32 variable from the
// process environment.
//
// Accepts a name of the variable.
//
// Returns the variable's value and an error if it exists,
// otherwise nil.
func GetFloatEnvVar(name string) (float32, error) {
	// get the var from the environment
	envStr := os.Getenv(name)
	if len(envStr) == 0 {
		return 0, errors.New(fmt.Sprintf("failed to get env variable: %s", name))
	}

	// parse string to int
	envInt, err := StringToFloat(envStr)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("failed to parse env variable: %s", name))
	}
	return envInt, nil
}
