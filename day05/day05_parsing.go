package day05

import (
	"regexp"
	"strconv"
	"strings"
)

type Container struct {
	Stacks       map[int][]string
	Instructions [][]int
}

var (
	MoveRegexp = regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)`)
)

// ParseInput and makes it into this weird container format and
// will return back a Container type. This type will have stacks of strings,
// and will have instructions in a list of a list of moves. 0 will be the position
// 1 will be the start and 2 will be the move too.
func ParseInput(input []string) (*Container, error) {

	var cntr = &Container{Instructions: make([][]int, 0)}

	for i := 0; i <= len(input)-1; i++ {
		l := input[i]

		var stacks = make(map[int][]string)

		if strings.HasPrefix(l, " 1   2   3") {
			for c := 0; c <= len(l)-1; c++ {
				char := string(l[c])
				if char == " " {
					continue
				}

				sn, err := strconv.Atoi(char)
				if err != nil {
					return nil, err
				}

				stacks[sn] = make([]string, 0)

				for j := i - 1; j >= 0; j-- {
					if len(input[j])-1 <= c {
						continue
					}

					char := string(input[j][c])
					if char == " " {
						continue
					}
					stacks[sn] = append(stacks[sn], char)
				}
			}

			cntr.Stacks = stacks
		}


		grps := MoveRegexp.FindStringSubmatch(l)
		if len(grps) == 0 {
			continue
		}

		start, err := strconv.Atoi(grps[1])
		if err != nil {
			return nil, err
		}

		from, err := strconv.Atoi(grps[2])
		if err != nil {
			return nil, err
		}

		to, err := strconv.Atoi(grps[3])
		if err != nil {
			return nil, err
		}

		cntr.Instructions = append(cntr.Instructions, []int{start, from, to})
	}

	return cntr, nil
}
