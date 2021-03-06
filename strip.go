package ansicode

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

var ansiRegexp = regexp.MustCompile("\x1b[^m]*m")

func stripAnsi(s string) string {
	return ansiRegexp.ReplaceAllLiteralString(s, "")
}

func stripReader(reader *bufio.Reader) {
	for {
		line, err := reader.ReadString('\n')
		if err == nil || err == io.EOF {
			fmt.Print(stripAnsi(line))
		}

		if err != nil {
			break
		}
	}
}

// StripStdin reads text from SDTIN and emits that same text minus any
// ANSI Escape codes.
func StripStdin() {
	reader := bufio.NewReader(os.Stdin)
	stripReader(reader)
}

// StripFile reads text from the file located at fileName and emits that same
// text minus any ANSI Escape codes.
func StripFile(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Error opening file %s, %v", fileName, err)
	}

	reader := bufio.NewReader(file)

	stripReader(reader)
}
