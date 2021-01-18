package ansicode

import (
	"fmt"
	"strconv"
	"strings"
)

// Color defines a custom color object which is defined by SGR parameters.
type Color struct {
	params []Attribute
}

// Attribute defines a single SGR Code
type Attribute int

const (
	Escape = "\x1b"
	End    = "\x1b[0m"
)

// Base attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity(Bright) text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

// New returns a newly created color object.
func New(value ...Attribute) *Color { return &Color{params: value} }

// sequence returns a formatted SGR sequence to be plugged into a "\x1b[...m"
// an example output might be: "1;36" -> bold cyan
func (c *Color) sequence() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

// Wrap wraps the s string with the colors attributes. The string is ready to be printed.
func (c Color) Wrap(s string) string { return c.Start() + s + End }

func (c Color) Start() string { return fmt.Sprintf("%s[%sm", Escape, c.sequence()) }

func (c Color) End() string { return End }

func (p Attribute) Start() string { return New(p).Start() }

func (p Attribute) End() string { return New(p).End() }

// Wrap wraps the s string with the colors attributes. The string is ready to be printed.
func (p Attribute) Wrap(s string) string { return New(p).Wrap(s) }
