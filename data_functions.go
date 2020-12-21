package dockerstats

import (
	"strings"
	"strconv"
)

// This funciton is used to take a list provided as a line seperated string and make an array from it
func stringToList(in string) ([]string, error) {
	parts := strings.Split(string(in), "\n")
	parts = parts[:len(parts)-1]

	// TODO: Better error handling
	return parts, nil
}

// This funciton parses a data string and separetes the number from the qualifier
func parseData(s string) (letters string, numbers float64, err error) {
	var l, n []rune
	for _, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
			l = append(l, r)
		case r >= 'a' && r <= 'z':
			l = append(l, r)
		case r == '.':
			n = append(n, r)
		case r >= '0' && r <= '9':
			n = append(n, r)
		}
	}
	nS := string(n)
	nF, err := strconv.ParseFloat(nS, 8)
	return string(l), nF, err
}