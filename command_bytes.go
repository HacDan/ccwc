package main

import "strconv"

func cmdBytes(b []byte) string {
	bytes := len(b)
	return strconv.Itoa(bytes)
}
