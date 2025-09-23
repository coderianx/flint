package flint

import (
	"fmt"
)

func Version() {
	const Version = "v1.4"
	fmt.Println("Flint Version:", Version)
	fmt.Println("https://github.com/coderianx/flint")
}

func Info() {
	const Version = "v1.4"
	const License = "MIT"
	const Author = "Coderian"
	const Release_date = "2025/09/24"
	fmt.Println("========== Flint Info ==========")
	fmt.Println("Flint Version:", Version)
	fmt.Println("License:", License)
	fmt.Println("Author:", Author)
	fmt.Println("Release_date:", Release_date)
	fmt.Println("===============================")

}
