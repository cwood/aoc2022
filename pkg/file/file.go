package file

import (
	"os"
	"strings"
)

func ReadLines(fp string) ([]string, error) {
	rawIn, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	var o []string
	strs := strings.Split(string(rawIn), "\n")

	for i, s := range strs {
		if s == "" && (i == 0 || i == len(strs)-1) {
			continue
		}
		o = append(o, s)
	}

	return o, nil
}
