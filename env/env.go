package env

import (
	"fmt"
	"os"
)

func GetEnv(name, fallback string) (result string) {
	defer fmt.Printf("%s=%s\n", name, result)
	if val, exists := os.LookupEnv(name); exists {
		result = val
		return
	}
	result = fallback
	return
}
