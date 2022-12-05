package parse

import "strings"

type LineGroup struct {
	Group []string
}

// ToLineGroup parses lines into groups and uses a empty line
// separator to designate a new group start
func ToLineGroup(lines []string) []LineGroup {

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
