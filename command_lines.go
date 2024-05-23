package main

import (
	"bytes"
	"strconv"
)

func cmdLines(b []byte) string {
	lineSep := []byte{'\n'}
	total := bytes.Count(b, lineSep)
	return strconv.Itoa(total)
}
