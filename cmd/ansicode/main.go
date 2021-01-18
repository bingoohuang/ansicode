package main

import (
	"fmt"
	"github.com/bingoohuang/ansicode"
)

func main() {
	// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797

	// 8-colors
	colors := []string{"Black", "Red", "Green", "Yellow", "Blue", "Magenta", "Cyan", "White"}

	fmt.Printf("Normal: ")
	for i := 30; i <= 37; i++ {
		fmt.Printf("\x1B[%dm[%d:%s]", i, i, colors[i-30])
	}
	fmt.Println("\u001B[0m")
	fmt.Printf("Normal: ")
	for i := 40; i <= 47; i++ {
		fmt.Printf("\x1B[%dm[%d:%s]", i, i, colors[i-40])
	}
	fmt.Println("\u001B[0m")
	fmt.Printf("Bright: ")
	for i := 90; i <= 97; i++ {
		fmt.Printf("\x1B[%dm[%d:%s]", i, i, colors[i-90])
	}
	fmt.Println("\u001B[0m")
	fmt.Printf("Bright: ")
	for i := 100; i <= 107; i++ {
		fmt.Printf("\x1B[%dm[%d:%s]", i, i, colors[i-100])
	}
	fmt.Println("\u001B[0m")

	// Most terminals, apart from the basic set of 8 colors, also support the "bright" or "bold" colors.
	// These have their own set of codes, mirroring the normal colors, but with an additional ;1 in their codes:
	// Set style to bold, red foreground.
	// \x1b[1;31mHello
	// Set style to dimmed white foreground with red background.
	// \x1b[2;37;41mWorld
	fmt.Printf("Bold  : ")
	for i := 30; i <= 37; i++ {
		fmt.Printf("\x1B[1;%dm[%d:%s]", i, i, colors[i-30])
	}
	fmt.Println("\u001B[0m")
	fmt.Printf("Bold  : ")
	for i := 40; i <= 47; i++ {
		fmt.Printf("\x1B[1;%dm[%d:%s]", i, i, colors[i-40])
	}
	fmt.Println("\u001B[0m")

	// 256 Colors
	// The following escape code tells the terminal to use the given color ID:
	// ESC[38;5;${ID}m

	// ESC Code Sequence	Description
	// ESC[38;5;${ID}m	Set foreground color.
	// ESC[48;5;${ID}m	Set background color.
	//fmt.Printf("256 Foreground: ")
	//for i := 0; i < 256; i++ {
	//	fmt.Printf("\x1B[38;5;%dm[%d]", i, i)
	//}
	//fmt.Println("\u001B[0m")
	//fmt.Printf("256 Background: ")
	//for i := 0; i < 256; i++ {
	//	fmt.Printf("\x1B[48;5;%dm[%d]", i, i)
	//}
	//fmt.Println("\u001B[0m")

	// RGB Colors
	// More modern terminals supports Truecolor (24-bit RGB), which allows you to set foreground and background colors using RGB.
	// These escape sequences are usually not well documented.
	// ESC Code Sequence	Description
	// ESC[38;2;{r};{g};{b}m	Set foreground color as RGB.
	// ESC[48;2;{r};{g};{b}m	Set background color as RGB.
	// Note that ;38 and ;48 corresponds to the 16 color sequence and is interpreted by the terminal
	// to set the foreground and background color respectively. Where as ;2 and ;5 sets the color format.

	fmt.Print(ansicode.FgCyan.Wrap("FgCyan."))
	fmt.Print(ansicode.FgBlue.Wrap("FgBlue."))
	fmt.Print(ansicode.FgRed.Wrap("FgRed."))
	fmt.Print(ansicode.FgMagenta.Wrap("FgMagenta."))
	fmt.Print(ansicode.New(ansicode.FgCyan, ansicode.Underline).Wrap("FgCyan/Underline."))
	fmt.Print(ansicode.New(ansicode.FgRed, ansicode.Bold).Wrap("FgRed/Bold."))
	fmt.Print(ansicode.New(ansicode.FgRed, ansicode.BgWhite).Wrap("FgRed/BgWhite."))
	fmt.Print(ansicode.New(ansicode.FgCyan, ansicode.Underline, ansicode.Bold).Wrap("FgCyan/Underline/Bold."))
	fmt.Print(ansicode.New(ansicode.FgRed, ansicode.Italic, ansicode.Concealed, ansicode.CrossedOut).Wrap("FgRed/Italic/Concealed/CrossedOut."))
	fmt.Println()
}
