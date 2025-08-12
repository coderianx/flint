package flint

import "fmt"

func Version() {
	const Version = "v1.0"
	fmt.Println("Flint Version:", Version)
	fmt.Println("https://github.com/grayvort3x/Flint")
}

func Info() {
	developer := "GrayVort3x"
	releaseDate := "08/2025"
	tiktok := "https://tiktok.com/@yazilim4313"
	github := "https://github.com/grayvort3x"
	telegram := "https://t.me/Grayvort3x"
	x_account := "https://x.com/grayvortex_dev"
	x_flint_account := "https://x.com/flintframework"
	github_flint := "https://github.com/grayvort3x/Flint"
	telegram_channel := "https://t.me/flint_framework"
	telegram_channel_tr := "https://t.me/flint_framework_tr"

	fmt.Println("┌────────────────────────────────────────────┐")
	fmt.Println("│                Flint Web Framework        │")
	fmt.Println("└────────────────────────────────────────────┘")
	fmt.Println("📅 Release Date  :", releaseDate)
	fmt.Println("👨‍💻 Developer    :", developer)
	fmt.Println("📢 Welcome to the Flint Web Framework!")
	fmt.Println()
	fmt.Println("🌐 Social Media Links")
	fmt.Println("   GitHub         :", github)
	fmt.Println("   TikTok         :", tiktok)
	fmt.Println("   Telegram       :", telegram)
	fmt.Println("   X Account      :", x_account)
	fmt.Println("──────────────────────────────────────────────")
	fmt.Println("🌐 Flint Links")
	fmt.Println("   GitHub         :", github_flint)
	fmt.Println("   Telegram Channel:", telegram_channel)
	fmt.Println("   Telegram Channel TR:", telegram_channel_tr)
	fmt.Println("   X Account          :", x_flint_account)
	fmt.Println("──────────────────────────────────────────────")
}
