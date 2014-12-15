// Package splash provides nice splash screens to display when starting a
// service.
package splash

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
)

// Splash contains all configuration and splash screens.
type Splash struct {
	MinWidth   int
	Foreground int
	Background int

	Splashes []string
}

// New returns a new Splash with default values set, but has no splashes.
func New() Splash {
	return Splash{
		MinWidth:   80,
		Foreground: 125,
		Background: 214,
	}
}

// AddSplash adds a splash. A splash is just a string, normally it consists of
// several lines separated by "\n".
func (s *Splash) AddSplash(splash string) {
	s.Splashes = append(s.Splashes, splash)
}

// WriteSplash writes a random splash screen to the given writer.
func (s Splash) WriteSplash(w io.Writer) {
	io.WriteString(w, s.pad(s.Splashes[rand.Intn(len(s.Splashes))]))
	io.WriteString(w, "\n")
}

// colorize wraps the provided string in some ANSI escape code sequences to
// make it all colorful and nice.
func (s Splash) colorize(str string) string {
	return fmt.Sprintf("\x1b[48;5;%vm\x1b[38;5;%vm\x1b[1m%v\x1b[0m", s.Background, s.Foreground, str)
}

// pad pads the given string to a fixed width by appenging whitespaces.
func (s Splash) pad(str string) string {
	lines := strings.Split(str, "\n")
	longest := s.MinWidth

	for _, line := range lines {
		if longest < len(line) {
			longest = len(line)
		}
	}

	padded := make([]string, len(lines))

	for i, line := range lines {
		paddedLine := []rune(strings.Repeat(" ", longest))
		for j, c := range line {
			paddedLine[j] = c
		}
		padded[i] = s.colorize(string(paddedLine))
	}

	return strings.Join(padded, "\n")
}
