package env

import (
	"fmt"
	"os"
)

func GetEnv(name, fallback string) string {
	if val, exists := os.LookupEnv(name); exists {
		fmt.Printf("%s = %s\n", name, val)
		return val
	}
	fmt.Printf("%s = %s\n", name, fallback)
	return fallback
}
