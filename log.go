package flint

import (
	"fmt"
)

func LogError(function string, message string) {
	fmt.Printf("\033[31m[Flint ERROR]\033[0m [%s] %s\n", function, message)
}
