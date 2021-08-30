package brutal

import (
	"bufio"
	"io"
	"os"
	"strings"

	"replaceText/rules"
)

// Not Graceful Method implement
type brutal struct {
}

func NewMethod() *brutal {
	return &brutal{}
}

func (b *brutal) Replace(path string, r rules.Rule) (err error) {
	var lineString string
	var eof error
	var same bool
	var newline string
	var pos = int64(0)

	fi, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)

	for {
		lineString, eof = br.ReadString('\n')
		if eof == io.EOF {
			break
		}

		same, newline = b.replaceLine(lineString, r)

		if same {
			pos += int64(len(newline + "\n"))

			continue
		}

		bytes := []byte(newline + "\n")
		if _, err = fi.WriteAt(bytes, pos); nil != err {
			return
		}

		pos += int64(len(newline + "\n"))
	}

	return
}

func (b *brutal) replaceLine(line string, r rules.Rule) (same bool, newLine string) {

	switch len(r.GetRegList()) {
	case 0:
		r.RangeRule(func(key, value string) bool {
			newLine = strings.Replace(line, key, value, -1)
			if newLine == line {
				same = true
			}
			return same
		})

	default:
		// todo:support reg
	}

	return
}
