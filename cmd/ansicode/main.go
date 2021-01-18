package main

import (
	"fmt"
	"github.com/bingoohuang/ansicode"
	"strconv"
)

func main() {
	const (
		InfoBoldColor    = "\033[1;34m%s\033[0m"
		NoticeBoldColor  = "\033[1;36m%s\033[0m"
		WarningBoldColor = "\033[1;33m%s\033[0m"
		ErrorBoldColor   = "\033[1;31m%s\033[0m"
		DebugBoldColor   = "\033[1;36m%s\033[0m"

		InfoColor    = "\033[34m%s\033[0m"
		NoticeColor  = "\033[36m%s\033[0m"
		WarningColor = "\033[33m%s\033[0m"
		ErrorColor   = "\033[31m%s\033[0m"
		DebugColor   = "\033[36m%s\033[0m"
	)

	fmt.Printf("Log\t:")
	fmt.Printf(InfoColor, "[Info]")
	fmt.Printf(NoticeColor, "[Notice]")
	fmt.Printf(WarningColor, "[Warning]")
	fmt.Printf(ErrorColor, "[Error]")
	fmt.Printf(DebugColor, "[Debug]")
	fmt.Println()
	fmt.Printf("Bold\t:")
	fmt.Printf(InfoBoldColor, "[Info]")
	fmt.Printf(NoticeBoldColor, "[Notice]")
	fmt.Printf(WarningBoldColor, "[Warning]")
	fmt.Printf(ErrorBoldColor, "[Error]")
	fmt.Printf(DebugBoldColor, "[Debug]")
	fmt.Println()

	fmt.Printf("Fg\t:")
	for i := ansicode.FgBlack; i <= ansicode.FgWhite; i++ {
		fmt.Printf(ansicode.New(i).Wrap("[" + strconv.Itoa(int(i)) + ":" + i.String() + "]"))
	}
	fmt.Println(ansicode.End)

	fmt.Printf("Bg\t:")
	for i := ansicode.BgBlack; i <= ansicode.BgWhite; i++ {
		fmt.Printf(ansicode.New(i).Wrap("[" + strconv.Itoa(int(i)) + ":" + i.String() + "]"))
	}
	fmt.Println(ansicode.End)

	for j := ansicode.CrossedOut; j <= ansicode.CrossedOut; j++ {
		fmt.Printf("Crossed:")
		for i := ansicode.FgBlack; i <= ansicode.FgWhite; i++ {
			c := ansicode.New(i, j)
			fmt.Printf(c.Wrap("[" + strconv.Itoa(int(i)) + ":" + i.String() + "]"))
		}
		fmt.Println(ansicode.End)
	}

	fmt.Printf("256\t:")
	const PrintColor = "\033[38;5;%dm[%03d]\033[39;49m"
	for j := 0; j < 256; j++ {
		if j > 0 && j%16 == 0 {
			fmt.Println()
			fmt.Print("\t ")
		}
		fmt.Printf(PrintColor, j, j)
	}
	fmt.Println()
}
