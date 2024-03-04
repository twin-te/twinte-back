package appenv

import (
	"fmt"
	"os"
	"strconv"
)

func loadBool(key string) bool {
	ret, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		panic(fmt.Errorf("failed to load environment variable (%s) due to %w", key, err))
	}
	return ret
}

func loadInt(key string) int {
	ret, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic(fmt.Errorf("failed to load environment variable (%s) due to %w", key, err))
	}
	return ret
}

func loadString(key string) string {
	ret := os.Getenv(key)
	if ret == "" {
		panic(fmt.Errorf("failed to load environment variable (%s), please set it", key))
	}
	return ret
}
