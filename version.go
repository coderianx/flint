package flint

import (
	"fmt"
)

func Version() {
	const Version = "v1.2.0"
	fmt.Println("Flint Version:", Version)
	fmt.Println("https://github.com/coderianx/flint")
}

func Info() {
	const Version = "v1.2.0"
	const License = "MIT"
	const Author = "CoderianX"
	const Release_date = "2025/08/12"
	fmt.Println("========== Flint Info ==========")
	fmt.Println("Flint Version:", Version)
	fmt.Println("License:", License)
	fmt.Println("Author:", Author)
	fmt.Println("Release_date:", Release_date)
	fmt.Println("===============================")

}
