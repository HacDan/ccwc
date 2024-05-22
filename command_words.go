package main

import (
	"strconv"
	"strings"
)

func cmdWords(b []byte) string {
	s := string(b)
	return strconv.Itoa(len(strings.Fields(s)))
}
