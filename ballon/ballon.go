package ballon

import (
	"bufio"
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/mitchellh/go-wordwrap"
	"log"
	"os"
	"strings"
)

var columns int32
var cows string

func init() {
	cows = `         \
          \
`
}

func ReadInput(args []string) []string {
	var tmps []string
	if len(args) == 0 {
		s := bufio.NewScanner(os.Stdin)

		for s.Scan() {
			tmps = append(tmps, s.Text())
		}

		if s.Err() != nil {
			log.Printf("failed reading stdin: %s\n", s.Err().Error())
			os.Exit(1)
		}

		if len(tmps) == 0{
			fmt.Println("Error: no input from stdin")
			os.Exit(1)
		}
	} else {
		tmps = args
	}

	var msgs []string
	for i := 0; i < len(tmps); i++ {
		expand := strings.Replace(tmps[i], "\t", "        ", -1)

		tmp := wordwrap.WrapString(expand, uint(columns))
		for _, s := range strings.Split(tmp, "\n") {
			msgs = append(msgs, s)
		}
	}
	return msgs
}

func MaxWidth(msgs []string) int {
	max := -1
	for _, m := range msgs {
		l := runewidth.StringWidth(m)
		if l > max {
			max = l
		}
	}

	return max
}

func SetPadding(msgs []string, width int) []string {
	var ret []string
	for _, m := range msgs {
		s := m + strings.Repeat(" ", width - runewidth.StringWidth(m))
		ret = append(ret, s)
	}

	return ret
}

func ConstructBallon(msgs []string, width int) string {
	var borders []string
	line := len(msgs)

	if line == 1 {
		borders = []string{"<", ">"}
	} else {
		borders = []string{"/", "\\", "\\", "/", "|", "|"}
	}

	var lines []string

	topBorder := " " + strings.Repeat("_", width+2)
	bottomBoder := " " + strings.Repeat("-", width+2)

	lines = append(lines, topBorder)
	if line == 1 {
		s := fmt.Sprintf("%s %s %s", borders[0], msgs[0], borders[1])
		lines = append(lines, s)
	} else {
		s := fmt.Sprintf(`%s %s %s`, borders[0], msgs[0], borders[1])
		lines = append(lines, s)
		i := 1
		for ; i < line-1; i++ {
			s = fmt.Sprintf(`%s %s %s`, borders[4], msgs[i], borders[5])
			lines = append(lines, s)
		}
		s = fmt.Sprintf(`%s %s %s`, borders[2], msgs[i], borders[3])
		lines = append(lines, s)
	}

	lines = append(lines, bottomBoder)
	return strings.Join(lines, "\n")
}
