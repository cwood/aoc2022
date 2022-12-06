package parse

import (
	"strconv"
	"strings"
)

type LineGroup struct {
	Group []string
}

// ToLineGroupWithCnt takes a input of strings and a group count and will
// create groups of line to that count
func ToLineGroupWithCnt(lines []string, c int) []LineGroup {

	var (
		groups = make([]LineGroup, 0)
		group  = LineGroup{Group: make([]string, 0)}
	)

	for _, l := range lines {
		if len(group.Group) == c {
			groups = append(groups, group)
			group = LineGroup{Group: make([]string, 0)}
		}
		group.Group = append(group.Group, l)
	}

	groups = append(groups, group)

	return groups
}

// ToLineGroupWithSep parses lines into groups and uses a empty line
// separator to designate a new group start
func ToLineGroupWithSep(lines []string) []LineGroup {

	var (
		groups = make([]LineGroup, 0)
		group  = LineGroup{Group: make([]string, 0)}
	)

	for _, l := range lines {
		if l == "" {
			groups = append(groups, group)
			group = LineGroup{Group: make([]string, 0)}
			continue
		}
		group.Group = append(group.Group, l)
	}
	groups = append(groups, group)

	return groups
}

// ToWords takes a string and splits it into words based on a space
func ToWords(lines []string) [][]string {

	var wl = make([][]string, 0)

	for _, l := range lines {
		w := strings.Split(l, " ")
		wl = append(wl, w)
	}

	return wl
}

// ToRanges takes lines with a int-int,int-int and will
// generate pairs of ints
func ToRanges(input []string) ([][][]int, error) {
	ranges := make([][][]int, 0)

	for _, r := range input {
		chunks := strings.Split(r, ",")
		grp := make([][]int, 0)
		for _, chnk := range chunks {
			tmns := strings.SplitN(chnk, "-", 2)
			s, err := strconv.Atoi(tmns[0])
			if err != nil {
				return nil, err
			}

			e, err := strconv.Atoi(tmns[1])
			if err != nil {
				return nil, err
			}

			grp = append(grp, []int{s, e})
		}

		ranges = append(ranges, grp)
	}

	return ranges, nil
}
